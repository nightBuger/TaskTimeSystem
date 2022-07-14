package webengine

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func CheckError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}

func CheckFatal(err error) {
	if err != nil {
		Logger.Fatal(err.Error())
		os.Exit(201)
	}
}

func CheckFmt(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(201)
	}
}

func trivalStructField(columns []string, structInfo reflect.Value) (correspond map[int]int) {
	correspond = make(map[int]int, 0)
loop:
	for i := 0; i < structInfo.NumField(); i++ {
		for j, column := range columns {
			if column == structInfo.Type().Field(i).Tag.Get("sql") {
				correspond[i] = j
				continue loop
			}
		}
	}
	return
}

func getStructPtr(structInfoPtr reflect.Value) (reflect.Value, error) {

	//check if ptr
	if structInfoPtr.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.New("must be a struct ptr")
	}
	//check if struct
	structInfo := structInfoPtr.Elem()
	if structInfo.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("must be a struct ptr")
	}

	return structInfo, nil
}
