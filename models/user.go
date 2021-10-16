package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (u *User) IsEmpty() bool {
	return u.FirstName == ""
}
