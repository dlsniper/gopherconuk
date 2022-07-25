package examples

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	socketio "github.com/googollee/go-socket.io"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SocketIoLogger(next func(w http.ResponseWriter, r *http.Request) *socketio.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

// setup routes
func (examples *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/posts", examples.Logger(examples.Post))
	mux.HandleFunc("/ws", examples.Logger(examples.MyWebSocket))

	mux.Handle("/socket.io", examples.SocketIoLogger(examples.MySocketIOServer))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
