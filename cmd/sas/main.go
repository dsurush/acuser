package main

import (
	"acuser/cmd/sas/app"
	"acuser/pkg/core/services"
	"acuser/pkg/core/token"
	"context"
	"flag"
	"github.com/dsurush/jwt/pkg/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
	dsn  = flag.String("dsn", "postgres://username:password@localhost:5500/app", "Postgres DSN")
)
func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)

	router := httprouter.New()
	pool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		log.Printf("%e", err)
	}

	svc := services.NewUserSvc(pool)
	tokenSvc := token.NewTokenSvc(svc, []byte(`surush`))
	secret := jwt.Secret(`surush`)
	server := app.NewMainServer(pool, router, svc, secret, tokenSvc)
	server.Start()
	
	panic(http.ListenAndServe(addr, server))
}