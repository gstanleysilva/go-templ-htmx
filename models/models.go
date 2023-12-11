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

// RemoveUser removes a user from the users slice
func (g *GlobalData) RemoveUser(id int) {
	for i, user := range g.Users {
		if user.ID == id {
			g.Users = append(g.Users[:i], g.Users[i+1:]...)
			break
		}
	}
}
