package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Id   int
	Name string
	Role int
}

type Role struct {
	Id     int
	Name   string
	Parent int
}

type UserRoleStructure struct {
	Roles []Role
	Users []User
}

func main() {
	userID := flag.Int("userID", 0, "User ID to get subordinate users of")
	rauf := flag.String("rolesAndUsersFile", "example.json", "Filepath to a JSON file describing the roles and users (refer to example.json for structure)")
	flag.Parse()
	urs, err := readUserRoleStructure(*rauf)

	if err != nil {
		panic(err)
	}
	u, err := urs.GetSubordinates(*userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Input User ID: %d\nInput Roles: %+v\nInput Users: \t\t %+v\nResulting Subordinates:\t %+v\n", *userID, urs.Roles, urs.Users, u)
}

// unsure if this is required or just used in the example for completeness
func (urs *UserRoleStructure) SetRoles(r []Role) {
	urs.Roles = []Role{}
	urs.Roles = append(urs.Roles, r...)
}

// unsure if this is required or just used in the example for completeness
func (urs *UserRoleStructure) SetUsers(u []User) {
	urs.Users = []User{}
	urs.Users = append(urs.Users, u...)
}

// this will be called to get all the subordinates for the specific UserRoleStructure
func (urs *UserRoleStructure) GetSubordinates(rId int) ([]User, error) {
	usersInit := make([]User, len(urs.Users))
	// take a copy so we dont impact the underlying slice
	_ = copy(usersInit, urs.Users)
	u, i := urs.getSubRolesRecurse(usersInit, rId)
	switch i {
	case 0:
		return u, nil
	case -1:
		return u, fmt.Errorf("input has no immediate parents")
	case -2:
		return u, fmt.Errorf("found some issues with the consistency of the role parent mappings")
	default:
		return u, fmt.Errorf("an unexpected error has occured")
	}
}

// need to pass in a full list of users and traverse up until 0 and remove entries until we get to the top
// assuming that there is no guarantee that anything will be sequential
func (urs *UserRoleStructure) getSubRolesRecurse(sr []User, rID int) ([]User, int) {
	// initialize to something out of bounds given that top level is 0
	newPar := -1
	srCopy := sr
	if rID == 0 || len(sr) == 0 {
		// we have reached the top, we can stop recursing
		// or we have exhausted everything and there are no subordinates
		return sr, 0
	}
	// lets find the parent to check
	for _, v := range urs.Roles {
		if v.Id == rID {
			newPar = v.Parent
		}
	}
	if newPar == -1 {
		// provided role provided doesn't exist
		// lets provide what we have
		if len(sr) == len(urs.Users) {
			// this means input was bad from the start
			return sr, -1
		} else {
			// this means that the relations between roles and parents is broken somewhere
			// may want to handle differently, for now just differentiating these cases
			return sr, -2
		}

	}
	for i := 0; i < len(srCopy); i++ {
		// if we find a user that matches the  input on the first iteration
		if len(sr) == len(urs.Users) && sr[i].Id == rID {
			srCopy = append(srCopy[:i], srCopy[i+1:]...)
		}
		// if we find a user that matches the parent role remove it
		if sr[i].Role == newPar {
			srCopy = append(srCopy[:i], srCopy[i+1:]...)
		}
	}
	// lets recurse and look up the tree
	return urs.getSubRolesRecurse(srCopy, newPar)
}

// just to reduce code duplication for testing but may be useful elsewhere
func isEqual(a, b UserRoleStructure) bool {
	if len(a.Roles) != len(b.Roles) {
		return false
	}
	for i := 0; i < len(a.Roles); i++ {
		if a.Roles[i] != b.Roles[i] {
			return false
		}
	}
	if len(a.Users) != len(b.Users) {
		return false
	}
	for i := 0; i < len(b.Users); i++ {
		if a.Users[i] != b.Users[i] {
			return false
		}
	}
	return true
}

// below just adding in to help run it in CLI and try a few different combinations of input
// not unit testing this explicitly because it is out of scope of the ask, just something useful to help you run it

func readUserRoleStructure(fp string) (*UserRoleStructure, error) {
	f, err := openFile(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	usr := &UserRoleStructure{}
	if err = json.Unmarshal(b, usr); err != nil {
		return nil, err
	}
	return usr, nil
}

func openFile(fp string) (*os.File, error) {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return nil, err
	}
	// read file contents
	return os.Open(fp)
}

// BELOW ONLY INCLUDED IN CASE CLI TOOL IS NOT OK FOR THE PURPOSES INTENDED
//
// init := &UserRoleStructure{}
// roles := []Role{
// 	{
// 		Id:     1,
// 		Name:   "System Administrator",
// 		Parent: 0,
// 	},
// 	{
// 		Id:     2,
// 		Name:   "Location Manager",
// 		Parent: 1,
// 	},
// 	{
// 		Id:     3,
// 		Name:   "Supervisor",
// 		Parent: 2,
// 	},
// 	{
// 		Id:     4,
// 		Name:   "Employee",
// 		Parent: 3,
// 	},
// 	{
// 		Id:     5,
// 		Name:   "Trainer",
// 		Parent: 3,
// 	},
// }
// users := []User{
// 	{
// 		Id:   1,
// 		Name: "Adam Admin",
// 		Role: 1,
// 	},
// 	{
// 		Id:   2,
// 		Name: "Emily Employee",
// 		Role: 4,
// 	},
// 	{
// 		Id:   3,
// 		Name: "Sam Supervisor",
// 		Role: 3,
// 	},
// 	{
// 		Id:   4,
// 		Name: "Mary Manager",
// 		Role: 2,
// 	},
// 	{
// 		Id:   5,
// 		Name: "Steve Trainer",
// 		Role: 5,
// 	},
// }
// init.SetRoles(roles)
// init.SetUsers(users)
// test := 3
// sub, err := init.GetSubordinates(test)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
