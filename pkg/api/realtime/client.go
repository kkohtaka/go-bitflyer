// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package realtime

import (
	jsonencoding "encoding/json"
	"fmt"
	"sync"
	"time"

	jsonrpc "github.com/gorilla/rpc/v2/json2"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/board"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/executions"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/ticker"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

const (
	APIEndpoint string = "wss://ws.lightstream.bitflyer.com/json-rpc"

	HeartbeatIntervalSecond time.Duration = 60
	ReadTimeoutSecond       time.Duration = 300
	WriteTimeoutSecond      time.Duration = 5
)

type Message struct {
	Version string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  ChannelMessage `json:"params"`
}

type ChannelMessage struct {
	Channel Channel                 `json:"channel"`
	Message jsonencoding.RawMessage `json:"message"`
}

type Channel string

var (
	ChannelOrderBook       Channel = "lightning_board_snapshot"
	ChannelOrderBookUpdate Channel = "lightning_board"
	ChannelTicker          Channel = "lightning_ticker"
	ChannelExecution       Channel = "lightning_executions"
)

type RequestParam struct {
	Channel Channel `json:"channel"`
}

// Client

type Client struct {
	Endpoint string
	Session  *Session
}

func NewClient() *Client {
	return &Client{
		Endpoint: APIEndpoint,
	}
}

func (c *Client) APIEndpoint() string {
	return c.Endpoint
}

func (c *Client) Connect() (*Session, error) {
	conn, _, err := websocket.DefaultDialer.Dial(c.Endpoint, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "open WebSocket connection to endpoint: %s", c.Endpoint)
	}
	return &Session{
		Conn: conn,
	}, nil
}

// Session

type Session struct {
	Conn *websocket.Conn
}

func (sess *Session) Close() error {
	err := sess.Conn.Close()
	if err != nil {
		return errors.Wrap(err, "close WebSocket session")
	}
	return nil
}

// Subscriber

type OrderBookHandler func(board.Response) error
type OrderBookUpdateHandler func(board.Response) error
type TickerHandler func(ticker.Response) error
type ExecutionHandler func(executions.Response) error

type Subscriber struct {
	mtx           sync.Mutex
	subscriptions []Channel
	handlers      map[Channel]interface{}
	msgc          chan ChannelMessage

	logger *zap.SugaredLogger
}

type SubscriberOpts struct {
	Debug bool
}

func NewSubscriber() *Subscriber {
	return NewSubscriberWithOpts(nil)
}

func NewSubscriberWithOpts(opts *SubscriberOpts) *Subscriber {
	var base *zap.Logger
	if opts == nil {
		opts = &SubscriberOpts{}
	}
	if opts.Debug {
		base, _ = zap.NewProduction()
	} else {
		base, _ = zap.NewDevelopment()
	}
	return &Subscriber{
		mtx:           sync.Mutex{},
		subscriptions: make([]Channel, 0),
		handlers:      map[Channel]interface{}{},
		msgc:          make(chan ChannelMessage),
		logger:        base.Sugar(),
	}
}

func (s *Subscriber) HandleOrderBook(
	pcs []markets.ProductCode,
	handler OrderBookHandler,
) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, pc := range pcs {
		ch := Channel(fmt.Sprintf("%s_%s", ChannelOrderBook, pc))
		s.subscriptions = append(s.subscriptions, ch)
		s.handlers[ch] = handler
	}
}

func (s *Subscriber) HandleOrderBookUpdate(
	pcs []markets.ProductCode,
	handler OrderBookUpdateHandler,
) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, pc := range pcs {
		ch := Channel(fmt.Sprintf("%s_%s", ChannelOrderBookUpdate, pc))
		s.subscriptions = append(s.subscriptions, ch)
		s.handlers[ch] = handler
	}
}

func (s *Subscriber) HandleTicker(
	pcs []markets.ProductCode,
	handler TickerHandler,
) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, pc := range pcs {
		ch := Channel(fmt.Sprintf("%s_%s", ChannelTicker, pc))
		s.subscriptions = append(s.subscriptions, ch)
		s.handlers[ch] = handler
	}
}

func (s *Subscriber) HandleExecution(
	pcs []markets.ProductCode,
	handler ExecutionHandler,
) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, pc := range pcs {
		ch := Channel(fmt.Sprintf("%s_%s", ChannelExecution, pc))
		s.subscriptions = append(s.subscriptions, ch)
		s.handlers[ch] = handler
	}
}

func (s *Subscriber) ListenAndServe(sess *Session) error {
	done := make(chan error)

	go func(sess *Session) {
		defer close(done)
		for {
			sess.Conn.SetReadDeadline(time.Now().Add(ReadTimeoutSecond * time.Second))
			_, data, err := sess.Conn.ReadMessage()
			if err != nil {
				done <- errors.Wrap(err, "read next RPC message")
				return
			}
			// TODO(kkohtaka): Handle responses of subscription requests
			var msg Message
			err = json.Unmarshal(data, &msg)
			if err != nil {
				s.logger.Warnw("unmarshal json",
					"err", err,
					"data", string(data),
				)
				continue
			}
			if msg.Params.Channel == "" {
				s.logger.Warnw("empty channel",
					"data", string(data),
				)
				continue
			}
			s.msgc <- msg.Params
		}
	}(sess)

	go func(sess *Session) {
		for msg := range s.msgc {
			func() {
				s.mtx.Lock()
				defer s.mtx.Unlock()

				handler, ok := s.handlers[msg.Channel]
				if ok {
					switch h := handler.(type) {
					case OrderBookHandler:
						var resp board.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							s.logger.Warnw("unmarshal json",
								"err", err,
								"channel", msg.Channel,
							)
						}
						err = h(resp)
						if err != nil {
							s.logger.Warnw("call handler",
								"err", err,
								"channel", msg.Channel,
							)
						}
					case OrderBookUpdateHandler:
						var resp board.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							s.logger.Warnw("unmarshal json",
								"err", err,
								"channel", msg.Channel,
							)
						}
						err = h(resp)
						if err != nil {
							s.logger.Warnw("call handler",
								"err", err,
								"channel", msg.Channel,
							)
						}
					case TickerHandler:
						var resp ticker.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							s.logger.Warnw("unmarshal json",
								"err", err,
								"channel", msg.Channel,
							)
						}
						err = h(resp)
						if err != nil {
							s.logger.Warnw("call handler",
								"err", err,
								"channel", msg.Channel,
							)
						}
					case ExecutionHandler:
						var resp executions.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							s.logger.Warnw("unmarshal json",
								"err", err,
								"channel", msg.Channel,
							)
						}
						err = h(resp)
						if err != nil {
							s.logger.Warnw("call handler",
								"err", err,
								"channel", msg.Channel,
							)
						}
					}
				} else {
					s.logger.Warnw("unknown channel",
						"channel", msg.Channel,
						"message", msg.Message,
					)
				}
			}()
		}
	}(sess)

	go func(sess *Session) {
		sess.Conn.SetPongHandler(func(string) error {
			s.logger.Debug("receive pong message")
			return nil
		})
		ticker := time.NewTicker(HeartbeatIntervalSecond * time.Second)
		for {
			select {
			case <-ticker.C:
				s.logger.Debug("send ping message")
				err := sess.Conn.WriteControl(
					websocket.PingMessage,
					[]byte{},
					time.Now().Add(WriteTimeoutSecond*time.Second),
				)
				if err != nil {
					ticker.Stop()
					sess.Close()
					return
				}
			}
		}
	}(sess)

	for _, channel := range s.subscriptions {
		param := RequestParam{
			Channel: channel,
		}
		msg, err := jsonrpc.EncodeClientRequest("subscribe", param)
		if err != nil {
			s.logger.Warnw("encode jsonrpc",
				"err", err,
				"data", param,
			)
			continue
		}
		sess.Conn.SetWriteDeadline(time.Now().Add(WriteTimeoutSecond * time.Second))
		err = sess.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			s.logger.Warnw("write message",
				"err", err,
			)
			continue
		}
	}

	return <-done
}
