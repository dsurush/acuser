package app

var(
	saveNewUser = `/save/NewClient`
	getClientByLogin = `/userList/:name`
	removeByLogin = `/userList/:name/remove`
	getUserList = `/userList`
)

func (server *MainServer) InitRouts() {
	server.router.POST(saveNewUser, server.SaveNewUserHandler)
	server.router.GET(getClientByLogin, server.GetUserByLoginHandler)
	server.router.DELETE(removeByLogin, server.RemoveUserByLoginHandler)
	server.router.GET(getUserList, server.GetUserListHandler)
}