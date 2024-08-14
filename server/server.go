package server

import "net/http"
import "log"
import "github.com/gorilla/handlers"

func NewServer(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger)

	var handler http.Handler = mux
	handler = handlers.CORS()(mux)
	handler = handlers.LoggingHandler(logger.Writer(), handler)
	return handler
}
