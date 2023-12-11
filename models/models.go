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

// AddUser adds a user to the users slice
func (g *GlobalData) AddUser(user User) {
	g.Users = append(g.Users, user)
}

// GetUser returns a user from the users slice based on the id provided. Returns nil if no user is found.
func (g *GlobalData) GetUser(id int) *User {
	for _, user := range g.Users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
