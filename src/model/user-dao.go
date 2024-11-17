package model

type UserDao struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Message     string `json:"message"`
}
