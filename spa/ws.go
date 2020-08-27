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
	c websocket.Conn
}

var (
	wserv wsServer
)

// ServeHTTP
func (ws wsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Warning Cors Header to '*'")
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:       []string{"echo"},
		InsecureSkipVerify: true, // Take care of CORS
		// OriginPatterns: ["*"],
	})
	if err != nil {
		log.Println("ERROR ", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "houston, we have a problem")

	log.Println("Wait a minute...")
	tQ := time.Tick(time.Second)

	cb := func(q Quote) {
		c1 := c
		q1 := q

		log.Println("ws [IN] message: ", q1)
		err = wsjson.Write(r.Context(), c1, q1)
		if err != nil {
			log.Println("ERROR: ", err)
		}
	}
	quoteCallbacks[c] = cb

	defer func() { delete(quoteCallbacks, c) }()
	func() {
		for {

			select {
			case now := <-tQ:
				t := NewTimeMsg(now)

				if config.Debug {
					log.Printf("Sending the time %+v", t)
				}
				err = wsjson.Write(r.Context(), c, t)
				if err != nil {
					log.Println("ERROR: ", err)
				}

			}
		}
	}()

	for {
		data := make([]byte, 8192)
		_, data, err := c.Read(r.Context())
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			log.Println("ws Closed")
			return
		}
		if err != nil {
			log.Println("ERROR: reading websocket ", err)
			return
		}
		log.Println("incoming: %s", string(data))
	}

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
