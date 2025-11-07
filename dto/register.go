package dto

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	RegisteredId     int `json:"registeredId"`
	RegisteredUserId int `json:"registeredUserId"`
}
