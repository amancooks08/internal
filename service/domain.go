package service

type RegisterUserRequest struct {
	Email string `json:"email"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Passwrod string `json:"password"`
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type LoginUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
}