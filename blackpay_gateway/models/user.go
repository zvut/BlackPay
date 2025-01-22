package models

type User struct {
	ID       string `json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Token    string `json:"token"`
	CSRF     string `json:"csrf"`
}
