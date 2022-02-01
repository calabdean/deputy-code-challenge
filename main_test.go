package main

import (
	"reflect"
	"testing"
)

const (
	errBadInput   = "input has no immediate parents"
	errRoleIssues = "found some issues with the consistency of the role parent mappings"
)

func TestGetSubordinates(t *testing.T) {
	init := &UserRoleStructure{}
	// lets do some testing with a basic object to ensure it works
	roles := []Role{
		{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		{
			Id:     2,
			Name:   "Location Manager",
			Parent: 1,
		},
		{
			Id:     3,
			Name:   "Supervisor",
			Parent: 2,
		},
		{
			Id:     4,
			Name:   "Employee",
			Parent: 3,
		},
		{
			Id:     5,
			Name:   "Trainer",
			Parent: 3,
		},
	}
	users := []User{
		{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		{
			Id:   2,
			Name: "Emily Employee",
			Role: 4,
		},
		{
			Id:   3,
			Name: "Sam Supervisor",
			Role: 3,
		},
		{
			Id:   4,
			Name: "Mary Manager",
			Role: 2,
		},
		{
			Id:   5,
			Name: "Steve Trainer",
			Role: 5,
		},
	}
	init.SetRoles(roles)
	init.SetUsers(users)
	tt := []struct {
		in      int
		init    UserRoleStructure
		wantU   []User
		wantErr string
		name    string
	}{
		{
			in:   3,
			init: *init,
			wantU: []User{
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
			wantErr: "",
			name:    "Basic check input 3",
		},
		{
			in:   1,
			init: *init,
			wantU: []User{
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
				{
					Id:   4,
					Name: "Mary Manager",
					Role: 2,
				},
				{
					Id:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
			wantErr: "",
			name:    "Basic check input 1",
		},
		{
			in:   0,
			init: *init,
			wantU: []User{
				{
					Id:   1,
					Name: "Adam Admin",
					Role: 1,
				},
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
				{
					Id:   4,
					Name: "Mary Manager",
					Role: 2,
				},
				{
					Id:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
			wantErr: "",
			name:    "Checking 0 parent",
		},
		{
			in:   -1,
			init: *init,
			wantU: []User{
				{
					Id:   1,
					Name: "Adam Admin",
					Role: 1,
				},
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
				{
					Id:   4,
					Name: "Mary Manager",
					Role: 2,
				},
				{
					Id:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
			wantErr: errBadInput,
			name:    "Checking negative input parent",
		},
		{
			in:   1,
			init: *init,
			wantU: []User{
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
				{
					Id:   4,
					Name: "Mary Manager",
					Role: 2,
				},
				{
					Id:   5,
					Name: "Steve Trainer",
					Role: 5,
				},
			},
			wantErr: "",
			name:    "Checking input of 1",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			subs, err := v.init.GetSubordinates(v.in)
			if err != nil {
				if err.Error() != v.wantErr {
					t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
				}
			}
			if !reflect.DeepEqual(v.wantU, subs) {
				t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
			}
		})
	}

	// lets do some testing with an empty object
	init.SetRoles([]Role{})
	init.SetUsers([]User{})
	tt = []struct {
		in      int
		init    UserRoleStructure
		wantU   []User
		wantErr string
		name    string
	}{
		{
			in:      1,
			init:    *init,
			wantU:   []User{},
			wantErr: errBadInput,
			name:    "Empty UserRoleStructure valid input",
		},
		{
			in:      -1,
			init:    *init,
			wantU:   []User{},
			wantErr: errBadInput,
			name:    "Empty UserRoleStructure negative input",
		},
		{
			in:      0,
			init:    *init,
			wantU:   []User{},
			wantErr: errBadInput,
			name:    "Empty UserRoleStructure 0 input",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			subs, err := v.init.GetSubordinates(v.in)
			if err != nil {
				if err.Error() != v.wantErr {
					t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
				}
			}
			if !reflect.DeepEqual(v.wantU, subs) {
				t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
			}
		})
	}

	// now lets look at potentially broken and unusual scenarios within the UserRoleStructure objects
	roles = []Role{
		{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		{
			Id:     2,
			Name:   "Location Manager",
			Parent: 4,
		},
		{
			Id:     3,
			Name:   "Supervisor",
			Parent: 1,
		},
	}
	users = []User{
		{
			Id:   1,
			Name: "Adam Admin",
			Role: 2,
		},
		{
			Id:   2,
			Name: "Emily Employee",
			Role: 4,
		},
		{
			Id:   3,
			Name: "Sam Supervisor",
			Role: 3,
		},
	}
	init.SetRoles(roles)
	init.SetUsers(users)
	tt = []struct {
		in      int
		init    UserRoleStructure
		wantU   []User
		wantErr string
		name    string
	}{
		{
			in:   3,
			init: *init,
			wantU: []User{
				{
					Id:   1,
					Name: "Adam Admin",
					Role: 2,
				},
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
			},
			wantErr: "",
			name:    "Check input valid and searches tree including disjoint user",
		},
		{
			in:   2,
			init: *init,
			wantU: []User{
				{
					Id:   1,
					Name: "Adam Admin",
					Role: 2,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
			},
			wantErr: errRoleIssues,
			name:    "Check input where user role mapping doesnt exist",
		},
		{
			in:   1,
			init: *init,
			wantU: []User{
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
				{
					Id:   3,
					Name: "Sam Supervisor",
					Role: 3,
				},
			},
			wantErr: errRoleIssues,
			name:    "Check scenario where user role parent doesnt exist",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			subs, err := v.init.GetSubordinates(v.in)
			if err != nil {
				if err.Error() != v.wantErr {
					t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
				}
			}
			if !reflect.DeepEqual(v.wantU, subs) {
				t.Errorf("\nGot %+v and %v\nWant %+v and %v", subs, err, v.wantU, v.wantErr)
			}
		})
	}
}

func TestSetRoles(t *testing.T) {
	tt := []struct {
		in   []Role
		want UserRoleStructure
		name string
	}{
		{
			in: []Role{
				{
					Id:     1,
					Name:   "System Administrator",
					Parent: 0,
				},
				{
					Id:     2,
					Name:   "Location Manager",
					Parent: 1,
				},
			},
			want: UserRoleStructure{
				Roles: []Role{
					{
						Id:     1,
						Name:   "System Administrator",
						Parent: 0,
					},
					{
						Id:     2,
						Name:   "Location Manager",
						Parent: 1,
					},
				},
				Users: []User{},
			},
			name: "Basic role loading",
		},
		{
			in: []Role{},
			want: UserRoleStructure{
				Roles: []Role{},
				Users: []User{},
			},
			name: "Empty role loading",
		},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			usr := &UserRoleStructure{}
			usr.SetRoles(v.in)
			if !isEqual(*usr, v.want) {
				t.Errorf("\nGot %+v\nWant %+v", *usr, v.want)
			}
		})
	}
}

func TestSetUsers(t *testing.T) {
	tt := []struct {
		in   []User
		want UserRoleStructure
		name string
	}{
		{
			in: []User{
				{
					Id:   1,
					Name: "Adam Admin",
					Role: 1,
				},
				{
					Id:   2,
					Name: "Emily Employee",
					Role: 4,
				},
			},
			want: UserRoleStructure{
				Users: []User{
					{
						Id:   1,
						Name: "Adam Admin",
						Role: 1,
					},
					{
						Id:   2,
						Name: "Emily Employee",
						Role: 4,
					},
				},
				Roles: []Role{},
			},
			name: "Basic user load",
		},
		{
			in: []User{},
			want: UserRoleStructure{
				Users: []User{},
				Roles: []Role{},
			},
			name: "Empty user load",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			usr := &UserRoleStructure{}
			usr.SetUsers(v.in)
			if !isEqual(*usr, v.want) {
				t.Errorf("\nGot %+v\nWant %+v", *usr, v.want)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	tt := []struct {
		inA  UserRoleStructure
		inB  UserRoleStructure
		want bool
		name string
	}{
		{
			inA: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			inB: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			want: true,
			name: "Basic equal",
		},
		{
			inA: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Wont Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			inB: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should not Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			want: false,
			name: "Roles different value",
		},
		{
			inA: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Wont Equal", Role: 1},
				},
			},
			inB: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Wont Equal above", Role: 1},
				},
			},
			want: false,
			name: "Users different value",
		},
		{
			inA: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			inB: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
					{Id: 2, Name: "Something else", Role: 2},
				},
			},
			want: false,
			name: "wrong length users",
		},
		{
			inA: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
					{Id: 2, Name: "Something else", Parent: 1},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			inB: UserRoleStructure{
				Roles: []Role{
					{Id: 1, Name: "Should Equal", Parent: 0},
				},
				Users: []User{
					{Id: 1, Name: "Should Equal", Role: 1},
				},
			},
			want: false,
			name: "wrong length roles",
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			got := isEqual(v.inA, v.inB)
			if v.want != got {
				t.Errorf("Want %v\nGot %v\n", v.want, got)
			}
		})
	}
}
