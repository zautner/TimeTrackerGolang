package services

import (
	"fmt"

	"test-time-tracker/data"
)

var (
	ErrorUser = fmt.Errorf("unknown user")
)

func AddUser(uname string) (e error) {
	e = data.CreateUser(uname)

	return
}
func DeleteUser(uname string) (e error) {
	e = data.DeleteUser(uname)
	return
}
func ListUsers() (answer interface{}, e error) {
	answer, e = data.ListUsers()
	return
}
