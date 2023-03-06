package test

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"seehuhn.de/go/websocket"
	"testing"
)

func TestWS(t *testing.T) {
	websocketHandler := &websocket.Handler{
		Handle: echo,
	}
	http.Handle("/", websocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(conn *websocket.Conn) {
	defer conn.Close(websocket.StatusOK, "")
	for {
		tp, r, err := conn.ReceiveMessage()
		if err != nil {
			break
		}
		w, err := conn.SendMessage(tp)
		if err != nil {
			// We still need to read the complete message, so that the
			// next read doesn't block.
			io.Copy(ioutil.Discard, r)
			break
		}
		_, err = io.Copy(w, r)
		if err != nil {
			io.Copy(ioutil.Discard, r)
		}
		w.Close()
	}
}
