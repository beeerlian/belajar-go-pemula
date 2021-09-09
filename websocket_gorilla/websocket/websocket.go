package websocket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	Upgrader.CheckOrigin = origin

	wsconn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return wsconn, nil
}
func origin(r *http.Request) bool {
	return true
}
func Writer(wsconn *websocket.Conn) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer func() {
		ticker.Stop()
		wsconn.Close()
	}()
	for tick := range ticker.C {
		if err := wsconn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("tick : %v", tick))); err != nil {
			return
		}
	}

}
