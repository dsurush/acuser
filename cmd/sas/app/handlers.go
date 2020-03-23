package app

import (
	"acuser/pkg/core/token"
	"acuser/pkg/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) SaveNewUserHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
	}
	return
}

func (server *MainServer) GetUserByLoginHandler(writer http.ResponseWriter, _ *http.Request, pr httprouter.Params) {
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

func (server *MainServer) RemoveUserByLoginHandler(writer http.ResponseWriter, _ *http.Request, pr httprouter.Params) {
	login := pr.ByName("name")
	err := server.userSvc.RemoveUserByLogin(login)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode("Deleted")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func (server *MainServer) GetUserListHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
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

func (server *MainServer) SaveNewUserDTOHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := server.userSvc.SaveUser(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	response := models.AddUserResponse{
		Error:       false,
		Description: "Пользователь добавлен",
		User:        user,
	}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}//
	return
}

func (server *MainServer) GetUsers(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, err := server.userSvc.GetUsers()
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

func (server *MainServer) CreateTokenHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		var requestBody token.RequestDTO
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
			log.Print(err)
			return
		}
//		log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
		response, err := server.tokenSVc.Generate(request.Context(), &requestBody)
		//log.Println(response)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
			if err != nil {
				log.Print(err)
			}
			return
		}
		cookie := http.Cookie{
			Name:     "token",
			Value:    response.Token,
			Expires:  time.Now().Add(time.Hour),
			HttpOnly: true,
			Path:     "/",
			// Domain:   "localhost",
		}
		http.SetCookie(writer, &cookie)
		err = json.NewEncoder(writer).Encode(&response)
		if err != nil {
			log.Print(err)
		}
}

func (server *MainServer) Test(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody token.RequestDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
} //test