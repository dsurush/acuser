package app

import (
	"acuser/pkg/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) SaveNewUserHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.userSvc.Save(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}// /-----> Возвращать что всё даанн )))
	return
}

func (server *MainServer) GetUserByLoginHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	login := pr.ByName("name")
	//log.Printf("%s\n", login)
	user, err := server.userSvc.GetUserByLogin(login)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}
func (server *MainServer) RemoveUserByLoginHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	login := pr.ByName("name")
//	log.Printf("%s\n", login)
	err := server.userSvc.RemoveUserByLogin(login)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}
func (server *MainServer) GetUserListHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	users, err := server.userSvc.GetUserList()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}