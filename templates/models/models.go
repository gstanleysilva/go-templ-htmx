package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GlobalData struct {
	Admin User   `json:"admin"`
	Users []User `json:"users"`
}
