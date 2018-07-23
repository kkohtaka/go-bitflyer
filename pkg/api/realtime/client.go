// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package realtime

import (
	jsonencoding "encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	jsonrpc "github.com/gorilla/rpc/v2/json2"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/board"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/executions"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/ticker"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

const (
	APIEndpoint string = "wss://ws.lightstream.bitflyer.com/json-rpc"
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
		return nil, err
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
	return sess.Conn.Close()
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
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		mtx:           sync.Mutex{},
		subscriptions: make([]Channel, 0),
		handlers:      map[Channel]interface{}{},
		msgc:          make(chan ChannelMessage),
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
			_, data, err := sess.Conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				done <- errors.New("session was closed")
				return
			}
			var msg Message
			err = json.Unmarshal(data, &msg)
			if err != nil {
				log.Println("unmarshal json:", err, string(data))
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
							log.Printf("unmarshal json: %s", err)
						}
						err = h(resp)
						if err != nil {
							log.Println(err)
						}
					case OrderBookUpdateHandler:
						var resp board.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							log.Printf("unmarshal json: %s", err)
						}
						err = h(resp)
						if err != nil {
							log.Println(err)
						}
					case TickerHandler:
						var resp ticker.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							log.Printf("unmarshal json: %s", err)
						}
						err = h(resp)
						if err != nil {
							log.Println(err)
						}
					case ExecutionHandler:
						var resp executions.Response
						err := json.Unmarshal(msg.Message, &resp)
						if err != nil {
							log.Printf("unmarshal json: %s", err)
						}
						err = h(resp)
						if err != nil {
							log.Println(err)
						}
					}
				} else {
					log.Println("unknown channel:", msg.Channel, msg.Message)
				}
			}()
		}
	}(sess)

	for _, channel := range s.subscriptions {
		param := RequestParam{
			Channel: channel,
		}
		msg, err := jsonrpc.EncodeClientRequest("subscribe", param)
		if err != nil {
			log.Fatal("encode:", err)
		}
		err = sess.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal("write message:", err)
		}
	}

	return <-done
}
