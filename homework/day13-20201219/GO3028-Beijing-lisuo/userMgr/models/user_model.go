package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

// UserQuit represent quit exit status code
const (
	AdminName string = "admin"
	AdminID   int64  = 1
	UserQuit  int    = 1
)

// User is user
type User struct {
	ID         int64  `orm:"column(id);pk;auto"`
	Name       string `orm:"size(100)"`
	Sex        int
	Address    string     `orm:"size(512)"`
	Cell       string     `orm:"size(512)"`
	Born       *time.Time `orm:"size(512)"`
	Password   string     `orm:"size(512)"`
	Created_at *time.Time `orm:"auto_now_add"`
	Updated_at *time.Time `orm:"auto_now"`
	Deleted_at *time.Time `orm:"null;"`
}

// UserList contains users
type UserList []User

// NewUser make a new user contains user's info
func NewUser(id int64, sex int, name, cell, address, passwd string, born *time.Time) User {
	return User{
		ID:       id,
		Name:     name,
		Sex:      sex,
		Cell:     cell,
		Address:  address,
		Born:     born,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
	}
}
