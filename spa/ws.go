package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type wsServer struct {
}

var (
	wserv wsServer
	wsQ   chan interface{}
)

func (ws wsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Take care of CORS
	log.Println("Warning Cors Header to '*'")
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:       []string{"echo"},
		InsecureSkipVerify: true,
		// OriginPatterns: ["*"],
	})
	if err != nil {
		log.Println("ERROR ", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "houston, we have a problem")

	// set up the channel and start listening
	wsQ = make(chan interface{})
	defer close(wsQ)

	log.Println("Wait a minute...")
	tQ := time.Tick(time.Second * 10)

	go func() {
		for {

			log.Printf("ws Getting our Select on [waiting for incoming ...] %d", len(wsQ))
			select {
			case now := <-tQ:
				t := NewTimeMsg(now)

				log.Printf("Sending the time %+v", t)
				err = wsjson.Write(r.Context(), c, t)
				if err != nil {
					log.Println("ERROR: ", err)
				}

			case msg := <-wsQ:
				log.Println("ws [IN] message: ", msg)
				err = wsjson.Write(r.Context(), c, msg)
				if err != nil {
					log.Println("ERROR: ", err)
				}
			}
		}
	}()

	for {
		data := make([]byte, 8192)
		_, data, err := c.Read(r.Context())
		if err != nil {
			log.Fatal("ERROR: reading websocket ", err)
		}
		log.Println("incoming: %s", string(data))
	}

}

func (ws *wsServer) SendMsg(msg interface{}) {
	if wsQ == nil {
		log.Fatal("Q has not been initialize")
	}
	wsQ <- msg
}

func echo(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	err = w.Close()
	return err
}
