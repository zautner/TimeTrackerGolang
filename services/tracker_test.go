package services

import (
	"fmt"
	"reflect"
	"testing"

	"test-time-tracker/models"
)

type argstate struct {
	uname string
	state models.UserState
}
type args struct {
	uname string
}

var (
	user1 = args{"user1"}
	user2 = args{"user2"}
	user3 = args{"user3"}
)

func TestAddUser(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good test1", user1, false,
		},
		{
			"Good test2", user2, false,
		},
		{
			"Good test3", user3, false,
		},
		{
			"Bad test2", user2, true,
		},
		{
			"Bad test1", user1, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUser(tt.args.uname); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good test2", user2, false,
		},
		{
			"Bad test2", user2, true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.uname); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestListUsers(t *testing.T) {
	tests := []struct {
		name       string
		wantAnswer map[string]string
		wantErr    bool
	}{
		{
			"List good", map[string]string{"user1": "Absent", "user3": "Absent"}, false,
		},
		{
			"List wrong", map[string]string{"user2": "Absent", "user3": "Absent"}, true,
		},
		{
			"List bad", map[string]string{"user1": "OOO", "user3": "Absent"}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnswer, err := ListUsers()
			if (err != nil || !reflect.DeepEqual(gotAnswer, tt.wantAnswer)) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{
				fmt.Printf("%+v\n", gotAnswer)
			}
		})
	}
}
func TestSetImplicitState(t *testing.T) {
	tests := []struct {
		name    string
		args    argstate
		wantErr bool
	}{
		{
			"Good test",
			argstate{
				uname: "user1",
				state: models.OOO,
			},
			false,
		},
		{
			"Good test1",
			argstate{
				uname: "user1",
				state: models.Working,
			},
			false,
		},
		{
			"Bad test",
			argstate{
				uname: "user2",
				state: models.Working,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetImplicitState(tt.args.uname, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("SetImplicitState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestThisMonth(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantD   []models.TrackerRecord
		wantErr bool
	}{
		{
			"Good test",
			user1,
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ThisMonth(tt.args.uname)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThisMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%+v\n", gotD)
		})
	}
}

func TestTrack(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Good user1", user1, false,
		},
		{
			"Good user3", user3, false,
		},
		{
			"Bad user2", user2, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Track(tt.args.uname); (err != nil) != tt.wantErr {
				t.Errorf("Track() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
