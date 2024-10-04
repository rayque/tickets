package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
