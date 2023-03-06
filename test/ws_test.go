// Copyright 2022-present The ZTDBP Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
