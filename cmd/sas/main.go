package main

import (
	"acuser/cmd/sas/app"
	"acuser/pkg/core/services"
	"context"
	"flag"
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
	server := app.NewMainServer(pool, router, svc)

	server.Start()
	
	http.ListenAndServe(addr, server)
}