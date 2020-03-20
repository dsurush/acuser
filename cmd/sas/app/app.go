package app

import (
	"acuser/pkg/core/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {
	pool *pgxpool.Pool
	router *httprouter.Router
	userSvc *services.UserSvc
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, userSvc *services.UserSvc) *MainServer {
	return &MainServer{pool: pool, router: router, userSvc: userSvc}
}

func (server *MainServer) Start() {
	err := server.userSvc.DbInit()
	if err != nil {
		panic("server don't created")
	}
	//login, err := server.userSvc.GetUserByLogin(`surush`)
	//	login, err := server.UserSvc.GetUserByLogin("surush")
	//log.Println(login.Name, login.Surname, login.Login, login.Address, login.Email, login.Phone, login.Role)
	server.InitRouts()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// delegation////
	server.router.ServeHTTP(writer, request)
}

