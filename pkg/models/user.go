package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Address  string `json:"address,omitempty"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Role     string  `json:"role,omitempty"`
	Remove   bool   `json:"remove,omitempty"`
}

type UserDTO struct {
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Login    string `json:"login,omitempty"`
	Address  string `json:"address,omitempty"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Role     string  `json:"role,omitempty"`
}

type AddUserResponse struct {
	Error       bool    `json:"error,omitempty"`
	Description string  `json:"description,omitempty"`
	User      UserDTO `json:"client"`
}

type RequestSurush struct {
	Login string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}