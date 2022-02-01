# Dean Balac Deputy Code Challenge Solution

## How to run
There is an example below on how to display the help options and following that, how to call the executable using the flags and values.

By default the data from the challenge data is populated, just need to provide the userID as input.

```console
dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ ./get-subordinates --help
Usage of /tmp/go-build1748857946/b001/exe/main:
  -rolesAndUsersFile string
        Filepath to a JSON file describing the roles and users (refer to example.json for structure) (default "example.json")
  -userID int
        User ID to get subordinate users of

dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ ./get-subordinates -userID 3 -rolesAndUsersFile example.json
Input User ID: 3
Input Roles: [{Id:1 Name:System Administrator Parent:0} {Id:2 Name:Location Manager Parent:1} {Id:3 Name:Supervisor Parent:2} {Id:4 Name:Employee Parent:3} {Id:5 Name:Trainer Parent:3}]
Input Users:             [{Id:1 Name:Adam Admin Role:1} {Id:2 Name:Emily Employee Role:4} {Id:3 Name:Sam Supervisor Role:3} {Id:4 Name:Mary Manager Role:2} {Id:5 Name:Steve Trainer Role:5}]
Resulting Subordinates:  [{Id:2 Name:Emily Employee Role:4} {Id:5 Name:Steve Trainer Role:5}]

dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ ./get-subordinates -userID 1
Input User ID: 1
Input Roles: [{Id:1 Name:System Administrator Parent:0} {Id:2 Name:Location Manager Parent:1} {Id:3 Name:Supervisor Parent:2} {Id:4 Name:Employee Parent:3} {Id:5 Name:Trainer Parent:3}]
Input Users:             [{Id:1 Name:Adam Admin Role:1} {Id:2 Name:Emily Employee Role:4} {Id:3 Name:Sam Supervisor Role:3} {Id:4 Name:Mary Manager Role:2} {Id:5 Name:Steve Trainer Role:5}]
Resulting Subordinates:  [{Id:2 Name:Emily Employee Role:4} {Id:3 Name:Sam Supervisor Role:3} {Id:4 Name:Mary Manager Role:2} {Id:5 Name:Steve Trainer Role:5}]
```

Please refer to example.json to see how to construct a User and Role JSON object that is taken as input for the program.
```json
{ 
    "roles": [
        {
            "Id": 1,
            "Name": "System Administrator",
            "Parent": 0
        },
        ...
    ],
   "users": [
        {
            "Id": 1,
            "Name": "Adam Admin",
            "Role": 1
        },
        ...
}
```

## How to test
The below shows how to run the tests.

The commands will also check the test coverage and also bring up a report in your browser when run.

To view the coverage after tests are run and the browser is closed you can run the command "go tool cover -html=cover.out"

```console
dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ make test
go test -v ./...
=== RUN   TestGetSubordinates
=== RUN   TestGetSubordinates/Basic_check_input_3
=== RUN   TestGetSubordinates/Basic_check_input_1
=== RUN   TestGetSubordinates/Checking_0_parent
=== RUN   TestGetSubordinates/Checking_negative_input_parent
=== RUN   TestGetSubordinates/Checking_input_of_1
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_valid_input
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_negative_input
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_0_input
=== RUN   TestGetSubordinates/Check_input_valid_and_searches_tree_including_disjoint_user
=== RUN   TestGetSubordinates/Check_input_where_user_role_mapping_doesnt_exist
=== RUN   TestGetSubordinates/Check_scenario_where_user_role_parent_doesnt_exist
--- PASS: TestGetSubordinates (0.00s)
    --- PASS: TestGetSubordinates/Basic_check_input_3 (0.00s)
    --- PASS: TestGetSubordinates/Basic_check_input_1 (0.00s)
    --- PASS: TestGetSubordinates/Checking_0_parent (0.00s)
    --- PASS: TestGetSubordinates/Checking_negative_input_parent (0.00s)
    --- PASS: TestGetSubordinates/Checking_input_of_1 (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_valid_input (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_negative_input (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_0_input (0.00s)
    --- PASS: TestGetSubordinates/Check_input_valid_and_searches_tree_including_disjoint_user (0.00s)
    --- PASS: TestGetSubordinates/Check_input_where_user_role_mapping_doesnt_exist (0.00s)
    --- PASS: TestGetSubordinates/Check_scenario_where_user_role_parent_doesnt_exist (0.00s)
=== RUN   TestSetRoles
=== RUN   TestSetRoles/Basic_role_loading
=== RUN   TestSetRoles/Empty_role_loading
--- PASS: TestSetRoles (0.00s)
    --- PASS: TestSetRoles/Basic_role_loading (0.00s)
    --- PASS: TestSetRoles/Empty_role_loading (0.00s)
=== RUN   TestSetUsers
=== RUN   TestSetUsers/Basic_user_load
=== RUN   TestSetUsers/Empty_user_load
--- PASS: TestSetUsers (0.00s)
    --- PASS: TestSetUsers/Basic_user_load (0.00s)
    --- PASS: TestSetUsers/Empty_user_load (0.00s)
=== RUN   TestIsEqual
=== RUN   TestIsEqual/Basic_equal
=== RUN   TestIsEqual/Roles_different_value
=== RUN   TestIsEqual/Users_different_value
=== RUN   TestIsEqual/wrong_length_users
=== RUN   TestIsEqual/wrong_length_roles
--- PASS: TestIsEqual (0.00s)
    --- PASS: TestIsEqual/Basic_equal (0.00s)
    --- PASS: TestIsEqual/Roles_different_value (0.00s)
    --- PASS: TestIsEqual/Users_different_value (0.00s)
    --- PASS: TestIsEqual/wrong_length_users (0.00s)
    --- PASS: TestIsEqual/wrong_length_roles (0.00s)
PASS
ok      deputy-code-challenge   0.011s
go test -coverprofile cover.out
PASS
coverage: 60.9% of statements
ok      deputy-code-challenge   0.009s
go tool cover -html=cover.out
```

## How to build
The below shows how to build the exectuable, this will build a file get-suborinates that is an exectuable which can be run.

The code was built on a linux environment so it should run perfectly fine on a linux machine. (compiled with CGO_ENABLED=0 so should further minimize chances of issues)

```console
dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ make clean build
rm get-subordinates
go build -o get-subordinates
dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ ls
example.json  get-subordinates  go.mod  main.go  main_test.go  makefile  README.md
```
## How to package for deployment
The below shows how to package all the source code and related files into a single deployment package.

Note that tests are run in order to create cover.out to be packaged also.

Also note that for a proper deployment, all that would be sent would be the executable and any config files required.

```console
dean@dean-VirtualBox:~/dev/go/deputy-code-challenge$ make package
rm -f get-subordinates dean-balac-deputy-code-challenge.tar.gz cover.out
go test -v ./...
=== RUN   TestGetSubordinates
=== RUN   TestGetSubordinates/Basic_check_input_3
=== RUN   TestGetSubordinates/Basic_check_input_1
=== RUN   TestGetSubordinates/Checking_0_parent
=== RUN   TestGetSubordinates/Checking_negative_input_parent
=== RUN   TestGetSubordinates/Checking_input_of_1
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_valid_input
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_negative_input
=== RUN   TestGetSubordinates/Empty_UserRoleStructure_0_input
=== RUN   TestGetSubordinates/Check_input_valid_and_searches_tree_including_disjoint_user
=== RUN   TestGetSubordinates/Check_input_where_user_role_mapping_doesnt_exist
=== RUN   TestGetSubordinates/Check_scenario_where_user_role_parent_doesnt_exist
--- PASS: TestGetSubordinates (0.00s)
    --- PASS: TestGetSubordinates/Basic_check_input_3 (0.00s)
    --- PASS: TestGetSubordinates/Basic_check_input_1 (0.00s)
    --- PASS: TestGetSubordinates/Checking_0_parent (0.00s)
    --- PASS: TestGetSubordinates/Checking_negative_input_parent (0.00s)
    --- PASS: TestGetSubordinates/Checking_input_of_1 (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_valid_input (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_negative_input (0.00s)
    --- PASS: TestGetSubordinates/Empty_UserRoleStructure_0_input (0.00s)
    --- PASS: TestGetSubordinates/Check_input_valid_and_searches_tree_including_disjoint_user (0.00s)
    --- PASS: TestGetSubordinates/Check_input_where_user_role_mapping_doesnt_exist (0.00s)
    --- PASS: TestGetSubordinates/Check_scenario_where_user_role_parent_doesnt_exist (0.00s)
=== RUN   TestSetRoles
=== RUN   TestSetRoles/Basic_role_loading
=== RUN   TestSetRoles/Empty_role_loading
--- PASS: TestSetRoles (0.00s)
    --- PASS: TestSetRoles/Basic_role_loading (0.00s)
    --- PASS: TestSetRoles/Empty_role_loading (0.00s)
=== RUN   TestSetUsers
=== RUN   TestSetUsers/Basic_user_load
=== RUN   TestSetUsers/Empty_user_load
--- PASS: TestSetUsers (0.00s)
    --- PASS: TestSetUsers/Basic_user_load (0.00s)
    --- PASS: TestSetUsers/Empty_user_load (0.00s)
=== RUN   TestIsEqual
=== RUN   TestIsEqual/Basic_equal
=== RUN   TestIsEqual/Roles_different_value
=== RUN   TestIsEqual/Users_different_value
=== RUN   TestIsEqual/wrong_length_users
=== RUN   TestIsEqual/wrong_length_roles
--- PASS: TestIsEqual (0.00s)
    --- PASS: TestIsEqual/Basic_equal (0.00s)
    --- PASS: TestIsEqual/Roles_different_value (0.00s)
    --- PASS: TestIsEqual/Users_different_value (0.00s)
    --- PASS: TestIsEqual/wrong_length_users (0.00s)
    --- PASS: TestIsEqual/wrong_length_roles (0.00s)
PASS
ok      deputy-code-challenge   (cached)
go test -coverprofile cover.out
PASS
coverage: 60.9% of statements
ok      deputy-code-challenge   0.010s
go tool cover -html=cover.out
export CGO_ENABLED=0
go build -o get-subordinates
tar -czvf dean-balac-deputy-code-challenge.tar.gz *
README.md
cover.out
example.json
get-subordinates
go.mod
main.go
main_test.go
makefile
```

## Explanation of method
In order to accomplish this I made a few assumptions:
* We could not rely on sequential role ownership (ie role 5 could have parent 0, and role 1 could be subordinate to role 5)
* The largest role may not be equal to the length of the role array

Using the above, the only way to find the subordinates efficiently by following the data of the roles and users is to traverse up the tree.
This only gets you the Superior users however.
To get the subordinate, I made a recursive function that would start at the noted user ID and remove all the superior users from the total user array.
This would leave the remainder which equates to the subordinate users. This function would stop recursing once it ends up on the top level role 0.

## Note
* I was unsure if it was ok to use a file to read in contents or if i was required to create the setters as in the documentation so I created them and wrote tests for it but it is alot easier and more flexible to set it as a JSON file for the data and create a CLI application rather than hard code values into code
  * For this reason I didn't write unit tests for the file reding and data marshalling
  * If this method was acceptable, I would have adopted this for unit test variables which make the test code a little long