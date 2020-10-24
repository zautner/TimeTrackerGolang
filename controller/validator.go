package controller

import (
	"test-time-tracker/data"
	"test-time-tracker/models"
)

func validateUser(uname string) {
	_, e := data.GetUserState(uname)
	if e != nil {
		panic(e)
	}
}
func validateState(state string) models.UserState {
	s, e := models.StringToState(state)
	if e != nil {
		panic(e)
	}
	return s
}

func validateCall(ret ...interface{}) interface{} {
	for _, i := range ret {
		if t, ok := i.(error); ok {
			if t != nil {
				panic(t)
			}
		}
	}
	return ret
}
