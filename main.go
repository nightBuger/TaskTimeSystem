package main

import (
	_ "TaskTimeSystem/service"
	. "TaskTimeSystem/webengine"
	"fmt"
)

func main() {
	fmt.Println("gogogo")

	GinInstance.Run()
}
