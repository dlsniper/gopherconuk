package examples

import (
	"errors"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	gorillaWebsocket "github.com/gorilla/websocket"
)

// Sends back ws message with slight transformation ("from GO SERVER")
type WSMessage struct {
	Message string `json:"message"`
}

var upgrader = gorillaWebsocket.Upgrader{}

func (h *Handlers) MyWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if !errors.Is(err, nil) {
		fmt.Println(err)
	}
	defer ws.Close()

	fmt.Println("Connected!")

	for {
		var message WSMessage
		err := ws.ReadJSON(&message)
		if !errors.Is(err, nil) {
			fmt.Printf("error occurred: %v", err)
			break
		}
		message.Message += " (from GO SERVER)"
		fmt.Println(message)
		// send message from server
		if err := ws.WriteJSON(message); !errors.Is(err, nil) {
			fmt.Printf("error occurred: %v", err)
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Web Socket Connected!"))
}
