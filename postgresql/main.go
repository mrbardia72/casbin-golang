package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cychiuae/casbin-pg-adapter"
)

func main() {
	db, err := connectionTODB()
	adapter := casbinpgadapterx(err, db)
	enforcer := configModel(err, adapter)

	err1 := enforcer.LoadPolicy()
	if err1 != nil {
		return
	}

	useCMD := flag.Int64("cmd1", 1, "display listRoles")
	useCMD = flag.Int64("cmd2", 2, "display addUserRole")
	useCMD = flag.Int64("cmd3", 3, "display deleteUserRole")
	useCMD = flag.Int64("cmd4", 4, "display createRolePermission")
	useCMD = flag.Int64("cmd5", 5, "display deleteRolePermission")
	flag.Parse()

	if *useCMD == 1 {
		listRoles(enforcer)
		return

	} else if *useCMD == 2 {
		addUserRole(enforcer)
		return

	} else if *useCMD == 3 {
		deleteUserRole(enforcer)
		return

	} else if *useCMD == 4 {
		createRolePermission(enforcer)
		return

	} else if *useCMD == 5 {
		deleteRolePermission(enforcer)
		return

	}

	err2 := enforcer.SavePolicy()
	if err2 != nil {
		return
	}
}

func listRoles(enforcer *casbin.Enforcer) {
	res := enforcer.GetAllRoles() //get all role type tag "g"
	fmt.Println(res)

}

func configModel(err error, adapter *casbinpgadapter.Adapter) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer("./model/model.conf", adapter)
	if err != nil {
		panic(err)
	}

	return enforcer
}

func casbinpgadapterx(err error, db *sql.DB) *casbinpgadapter.Adapter {
	tableName := "casbin"
	adapter, err := casbinpgadapter.NewAdapter(db, tableName)
	if err != nil {
		panic(err)
	}

	return adapter
}

func connectionTODB() (*sql.DB, error) {
	connectionString := "host=127.0.0.1 user=admin password=adminpw dbname=casbin sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return db, err
}

func createRolePermission(enforcer *casbin.Enforcer) {
	bol, _ := enforcer.AddPolicy("bardia", "post2", "write4")
	if bol {
		fmt.Println("yes createRolePermission")
	} else {
		fmt.Println("no createRolePermission")
	}
}

func addUserRole(enforcer *casbin.Enforcer) {
	bol, _ := enforcer.AddGroupingPolicy("sina", "post")
	if bol {
		fmt.Println("yes addUserRole")
	} else {
		fmt.Println("no addUserRole")
	}
}

func deleteRolePermission(enforcer *casbin.Enforcer) {
	bol, _ := enforcer.RemovePolicy("alice", "data2", "write")
	if bol {
		fmt.Println("yes deleteRolePermission")
	} else {
		fmt.Println("no deleteRolePermission")
	}
}

func deleteUserRole(enforcer *casbin.Enforcer) {
	bol, _ := enforcer.RemoveGroupingPolicy("sina", "data56_admin_month_30")
	if bol {
		fmt.Println("yes deleteUserRole")
	} else {
		fmt.Println("no deleteUserRole")
	}
}
