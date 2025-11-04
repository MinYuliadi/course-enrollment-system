package dto

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	RegisteredId int `json:"registeredId"`
}
