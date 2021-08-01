package main

import (
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"strings"
)

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	bardiaRole(e)
	erfanRole()
}

func erfanRole() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	erfanRoles, _ := e.GetRolesForUser("erfan")

	erfanPerms, _ := e.GetImplicitPermissionsForUser("erfan")

	fmt.Printf(
		"erfan roles: %v\n erfan permissions: %v\n",
		erfanRoles,
		erfanPerms,
	)
}

type myPer struct {
	Subject string
	Role    string
	Per     [][]string
}

func bardiaRole(e *casbin.Enforcer) {
	name := e.GetAllSubjects() //[role:admin role:user]

	bardiaRoles, _ := e.GetRolesForUser("bardia") //[role:admin]

	bardiaPerms, _ := e.GetImplicitPermissionsForUser("bardia") // [[role:admin data1 read] [role:admin data2 write]]

	checkPolicy(e)

	getUsersForRole, _ := e.GetUsersForRole("role:admin") //[bardia]

	HasRoleForUser, _ := e.HasRoleForUser("bardia", "role:admin") //true

	addRoleForUser(e)

	addPermissionForUser(e)

	fmt.Printf(
		"bardia roles: %v\n bardia permissions : %v\n all subject:%v\n GetUsersForRole:%v\n HasRoleForUser:%v\n",
		bardiaRoles,
		bardiaPerms,
		name,
		getUsersForRole,
		HasRoleForUser,
	)
	ss := strings.Join(getUsersForRole, ", ")
	ssf := strings.Join(bardiaRoles, ", ")
	jsondat := &myPer{Subject: ss, Role: ssf, Per: bardiaPerms}
	rolesubject, _ := json.MarshalIndent(jsondat, "", "    ")
	fmt.Println(string(rolesubject))
}

func addPermissionForUser(e *casbin.Enforcer) {
	bol1, _ := e.AddPermissionForUser("bardiax", "read")

	if bol1 {
		fmt.Println("creating success")
	} else {
		fmt.Println("creating fail")
	}
}

func addRoleForUser(e *casbin.Enforcer) {
	bol, _ := e.AddRoleForUser("bardiax", "data2_admin")

	if bol {
		fmt.Println("creating success")
	} else {
		fmt.Println("creating fail")
	}
}

func checkPolicy(e *casbin.Enforcer) {
	hasPolicy := e.HasPolicy("role:user", "data", "read")
	if hasPolicy {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
