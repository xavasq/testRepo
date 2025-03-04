package models

type User struct {
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}
