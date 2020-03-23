package app

import (
	"acuser/middleware/jwt"
	"acuser/middleware/logger"
	"acuser/pkg/core/token"
	"reflect"
)

/// /api/clients/
/// GET
///
const saveNewUser = `/api/users/save`
const getClientByLogin = `/api/users/:name`
const  removeByLogin = `/api/users/:name/remove`
const  getUserList = `/api/users`

func (server *MainServer) InitRouts() {
	handler := jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), server.secret)(server.RemoveUserByLoginHandler)

	server.router.POST(saveNewUser, logger.Logger("Save User")(server.SaveNewUserDTOHandler))
	server.router.GET(getClientByLogin, logger.Logger("Get Client By login")(server.GetUserByLoginHandler))
	server.router.GET(getUserList, logger.Logger("Get Users")(server.GetUsers))
	server.router.DELETE(removeByLogin, handler)
	server.router.POST(`/api/tokens`, logger.Logger("Get Token")(server.CreateTokenHandler))
//	server.router.DELETE(removeByLogin, server.RemoveUserByLoginHandler)
//	server.router.GET(getUserList, server.GetUserListHandler)
//	server.router.POST(saveNewUser, server.SaveNewUserHandler)
}