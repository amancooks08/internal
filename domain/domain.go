package domain

type RegisterUserRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Passwrod    string `json:"password"`
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
}
type User struct {
	ID          int64  `db:"id" json:"id"`
	Email       string `db:"email" json:"email"`
	Name        string `db:"name" json:"name"`
	PhoneNumber string `db:"number" json:"phone_number"`
	Password    string `db:"password" json:"password"` // Don't return password
}
