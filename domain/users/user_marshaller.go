package users

import (
	"encoding/json"
)

// PublicUser type
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser type
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall func
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		// if the keys of the json structures differ from each other
		// your only option is to map the fields individually
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// If the fields don't differ between structs (User and PrivateUser), you can simply use Marshall
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)

	return privateUser
}

// Marshall func Users
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}
	return result
}
