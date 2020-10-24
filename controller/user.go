package controller

import (
	"test-time-tracker/services"
)

func AddUser(uname string) {
	validateCall(services.AddUser(uname))
}

func SetUserState(uname string, state string) {
	validateUser(uname)
	s := validateState(state)

	validateCall(services.SetImplicitState(uname, s))
}
func DeleteUser(uname string) {
	validateUser(uname)
	validateCall(services.DeleteUser(uname))
}
func ListUsers() (list interface{}) {
	var err error
	list, err = services.ListUsers()
	validateCall(list, err)
	return
}
