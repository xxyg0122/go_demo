package tests

import (
	"fmt"
	"go_demo/dbManager"
)

func Insert()  {
	result,err:=dbManager.ExecCommand("insert into user(name) value('aafe')" )
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
	rows,_:=dbManager.Query("select * from user w")
	var user user
	for rows.Next() {
		rows.Scan(&user.id, &user.name);
		fmt.Println(user);
	}
}

type user struct {
	id int
	name string
}

func BatchInsert(){
	executors := make([]dbManager.Executor, 0)
	executors = append(executors,   dbManager.ToExecutor("insert into user(name) value(?)" ,"batchName1"))
	executors = append(executors, dbManager.ToExecutor("insert into user(name) value(?)" ,"batchName2"))

	dbManager.Execs(executors...)
}