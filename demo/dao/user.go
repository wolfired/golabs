package dao

import (
	"github.com/satori/go.uuid"
)

type User struct {
	UUID     uuid.UUID
	Username string
	Password string
	Email    string
}

func (u *User) Save() {

}
