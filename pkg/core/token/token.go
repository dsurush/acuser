package token

import (
	"acuser/pkg/core/services"
	"context"
	"errors"
	"github.com/dsurush/jwt/pkg/jwt"

	//"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type TokenSvc struct {
	UserSvc *services.UserSvc
	secret []byte
}

func NewTokenSvc(userSvc *services.UserSvc, secret []byte) *TokenSvc {
	return &TokenSvc{UserSvc: userSvc, secret: secret}
}



type Payload struct {
	Id    int64    `json:"id"`
	Exp   int64    `json:"exp"`
	Role string    `json:"role"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token string `json:"token"`
}

var ErrInvalidLogin = errors.New("invalid password")
var ErrInvalidPassword = errors.New("invalid password")


func (receiver *TokenSvc) Generate(context context.Context, request *RequestDTO) (response ResponseDTO, err error) {
	login, err := receiver.UserSvc.GetUserByLogin(request.Username)
	if err != nil {
		log.Printf("%e", ErrInvalidLogin)
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(request.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		err = ErrInvalidPassword
		return
	}

	response.Token, err = jwt.Encode(Payload{
		Id:    login.Id,
		Exp:   time.Now().Add(time.Hour).Unix(),
		Role:  login.Role,
	}, receiver.secret)
	return
}
