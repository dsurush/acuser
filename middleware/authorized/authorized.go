package authorized

import (
	"acuser/pkg/core/token"
	"context"
	"log"
	"net/http"
)

func Authorized(role string, payload func(ctx context.Context) interface{}) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			auth := payload(request.Context()).(*token.Payload)
				if role == `admin` || role == `student` {
					log.Printf("access granted %v %v", role, auth)
					next(writer, request)
					return
				}
			http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}
}

