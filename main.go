package main

import (
	"TaskTimeSystem/service"
	_ "TaskTimeSystem/service"
	. "TaskTimeSystem/webengine"
	"fmt"
)

func main() {
	fmt.Println("gogogo")
	var u service.UserInfoGetRes
	Query("select * from t_userb", &u)
	GinInstance.Run()
}
