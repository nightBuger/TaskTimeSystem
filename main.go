package main

import (
	_ "TaskTimeSystem/service"
	. "TaskTimeSystem/webengine"
	"fmt"
)

func main() {
	fmt.Println("gogogo")
	//var u service.UserInfoGetRes
	//Query("select userid,username,guid,secretlevel,email,lxdh from t_userb", &u)
	//Query("select * from t_userb", &u)

	GinInstance.Run()
}
