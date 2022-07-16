package main

import (
	_ "TaskTimeSystem/service"
	. "TaskTimeSystem/webengine"
	"fmt"
)

func main() {
	fmt.Println("gogogo")
	str := fmt.Sprintf("%s %s", "123", 123)
	fmt.Println(str)
	GinInstance.Run()
}
