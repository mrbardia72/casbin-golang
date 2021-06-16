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
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)

	ss := memory.NewAdapter()

	e, _ := casbin.NewEnforcer(m, ss)

	_, err := e.AddPolicy("bardia", "chaincode", "read")
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
}
