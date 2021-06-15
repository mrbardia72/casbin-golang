package main

import (
	"casbin-sample/adapter/bardiaadapter"
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

	b := []byte{}

	//var s bardiaadapter.Adapter(&b)

	ss := bardiaadapter.NewAdapter(&b)
	e, _ := casbin.NewEnforcer(m, ss)

	_, err := e.AddPolicy()
	if err != nil {
		return
	}

	_, err = e.RemovePolicy()
	if err != nil {
		return
	}

	err = e.SavePolicy()
	if err != nil {
		return
	}

	bol, _ := e.AddRoleForUser("golang", "data2admin")

	if bol {
		fmt.Println("creating success")
	} else {
		fmt.Println("creating fail")
	}
}
