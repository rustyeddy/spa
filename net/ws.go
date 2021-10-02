package net

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
)

type WebSocket struct {
	websocket.Conn
}

// ServeHTTP
func (ws WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Warning Cors Header to '*'")
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:       []string{"echo"},
		InsecureSkipVerify: true, // Take care of CORS
	})

	if err != nil {
		log.Println("ERROR ", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "houston, we have a problem")

	log.Println("Wait a minute...")
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
