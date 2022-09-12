package model

type User struct {
	ID        uint64 `json:"Id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ModelName string `json:"modelName"`
}
