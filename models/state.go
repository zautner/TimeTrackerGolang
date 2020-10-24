package models

import (
	"errors"
	"strings"
)

type UserState byte

const (
	Absent UserState = iota
	Working
	Sick
	ReserveDuty
	OOO
	labels  = "Absent,Working,Sick,ReserveDuty,OOO"
	unknown = "unknown user state"
)

var (
	StateLables  = strings.Split(labels, ",")
	ErrorUnknown = errors.New(unknown)
)

func (us UserState) String() string {
	return StateLables[us]
}
func StringToState(l string) (s UserState, e error) {
	for k, a := range StateLables {
		if a == l {
			s = UserState(k)
			return
		}
	}
	s, e = Absent, ErrorUnknown
	return
}
