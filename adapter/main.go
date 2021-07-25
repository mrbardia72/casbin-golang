package main

import (
	"casbin-sample/adapter/memory"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

func main() {
	m, _ := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)

	ss := memory.NewAdapter()

	e, _ := casbin.NewEnforcer(m, ss)

	_, err := e.AddPolicy("erfan", "asp", "execute")
	if err != nil {
		panic(fmt.Errorf("error on add policy: %w", err))
	}

	_, err = e.AddPolicy("bardia", "chaincode", "read")
	if err != nil {
		panic(fmt.Errorf("error on add policy: %w", err))
	}

	_, err = e.AddPolicy("naser", "chaincode", "read")
	if err != nil {
		panic(fmt.Errorf("error on add policy: %w", err))
	}

	_, err = e.AddPolicy("naser", "chaincode", "write")
	if err != nil {
		panic(fmt.Errorf("error on add policy: %w", err))
	}

	// check 1
	enforce, err := e.Enforce("bardia", "chaincode", "read")
	if err != nil {
		panic(fmt.Errorf("error on enforce: %w", err))
	}

	fmt.Printf("bardia can read chaincode? %v\n", enforce)

	// check 2
	enforce, err = e.Enforce("bardia", "chaincode", "write")
	if err != nil {
		panic(fmt.Errorf("error on enforce: %w", err))
	}

	fmt.Printf("bardia can write chaincode? %v\n", enforce)

	// check 3
	enforce, err = e.Enforce("naser", "chaincode", "write")
	if err != nil {
		panic(fmt.Errorf("error on enforce: %w", err))
	}

	fmt.Printf("naser can write chaincode? %v\n", enforce)

	_, err = e.RemovePolicy("naser", "chaincode", "write")

	// check 3
	enforce, err = e.Enforce("naser", "chaincode", "write")
	if err != nil {
		panic(fmt.Errorf("error on enforce: %w", err))
	}

	fmt.Printf("naser can write chaincode? %v\n", enforce)

	fmt.Println("---------------------------------")

	naserPerms, _ := e.GetImplicitPermissionsForUser("naser")
	fmt.Println("naserPerms", naserPerms)

	name := e.GetAllSubjects()
	fmt.Println("GetAllSubjects", name)

	allNamedObjects := e.GetAllNamedObjects("p")
	fmt.Println("allNamedObjects", allNamedObjects)

	allActions := e.GetAllActions()
	fmt.Println("allActions", allActions)

	bol, _ := e.AddRoleForUser("bardiax", "data2_admin")

	if bol {
		fmt.Println("creating success")
	} else {
		fmt.Println("creating fail")
	}

	allRoles := e.GetAllRoles()
	fmt.Println(allRoles)

	bardiaRoles, _ := e.GetRolesForUser("bardiax") //[role:admin]

	fmt.Println(bardiaRoles)

}
