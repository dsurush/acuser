package logger

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Logger(prefix string) func(next httprouter.Handle,) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
			log.Printf(
				"%s Method: %s, path: %s",
				prefix,
				request.Method,
				request.URL.Path,
			)
			next(writer, request, pr)
		}
	}
}
