package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

type wsServer struct {
}

func (ws wsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"echo"},
	})
	if err != nil {
		log.Println("ERROR ", err)
		return
	}

	defer c.Close(websocket.StatusInternalError, "houston, we have a problem")

	if c.Subprotocol() == "echo" {
		//c.Close(websocket.StatusPolicyViolation, "we are demanding the client speak echo!")
		fmt.Println("The client speaks ECHO!!!")
	}

	for {
		err = echo(r.Context(), c)
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
