package types

import "encoding/json"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) String() string {
	user, _ := json.Marshal(u)
	return string(user)
}

func CreateExampleUser() User {
	return User{
		ID:   0,
		Name: "Райн Гослинг",
		Age:  52,
	}
}
