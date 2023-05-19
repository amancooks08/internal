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
	Token 	string `json:"token"`
}
type User struct {
	ID          int64  `db:"id" json:"id"`
	Email       string `db:"email" json:"email"`
	Name        string `db:"name" json:"name"`
	PhoneNumber string `db:"number" json:"phone_number"`
	Password    string `db:"password" json:"password"` // Don't return password
}

type LoginDbResponse struct {
	ID  	  int64  `db:"id" json:"id"`
	Password	string `db:"password" json:"-"`
}

type Wallet struct {
	ID      int64   `db:"id" json:"id"`
	UserID  int64   `db:"user_id" json:"-"`
	Balance float64 `db:"balance" json:"balance"`
	CreationDate string `db:"creation_date" json:"creation_date"`
	LastUpdated string `db:"last_updated" json:"last_updated"`
	Status string `db:"status" json:"status"`
}

type GetWalletResponse struct {
	ID      int64   `db:"id" json:"id"`
	Balance float64 `db:"balance" json:"balance"`
	CreationDate string `db:"creation_date" json:"creation_date"`
	LastUpdated string `db:"last_updated" json:"last_updated"`
	Status string `db:"status" json:"status"`
}

type Message struct {
	Message string `json:"message"`
}

type Credit struct {
	Amount float64 `json:"amount"`
}