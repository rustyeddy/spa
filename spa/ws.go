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
	Q chan string
}

// NewWsServer creates a new webssocket server
func newWsServer() (ws *wsServer) {
	ws = &wsServer{}
	return ws
}

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

	log.Println("Wait a minute...")
	for now := range time.Tick(time.Second * 10) {
		t := NewTimeMsg(now)

		log.Printf("Sending the time %+v", t)
		err = wsjson.Write(r.Context(), c, t)
		if err != nil {
			log.Println("ERROR: ", err)
		}
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
