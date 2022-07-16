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
	for structIndex := 0; structIndex < structInfo.NumField(); structIndex++ {
		for columnIndex, column := range columns {
			if column == structInfo.Type().Field(structIndex).Tag.Get("sql") {
				correspond[structIndex] = columnIndex
				continue loop
			}
		}
	}
	return
}

func bindInterfaceSlice(interfacePtr *[]interface{}, structInfo reflect.Value, corresopnd map[int]int) {
	for structIndex, columnIndex := range corresopnd {
		(*interfacePtr)[columnIndex] = structInfo.Elem().Field(structIndex).Addr().Interface()
	}
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

func mustBeStruct(value reflect.Value) bool {
	return value.Kind() == reflect.Struct
}
