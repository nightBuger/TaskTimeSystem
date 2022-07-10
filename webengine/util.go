package webengine

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}

func CheckFatal(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
	os.Exit(201)
}

func CheckFmt(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(201)
	}
}
