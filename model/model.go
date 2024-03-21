package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BasicResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
