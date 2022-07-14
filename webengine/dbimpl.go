package webengine

import (
	"errors"
	"reflect"
)

func QueryRow(sqlstr string, res interface{}) (err error) {
	//检查是否为指针
	ptr := reflect.ValueOf(res)
	if ptr.Kind() != reflect.Ptr {
		Logger.Error("sqlstr:[", sqlstr, "]的传入res类型不正确,传入的res的类型为:", ptr.Kind(), ",实际上应当为指针类型")
		err = errors.New("res应当为指针类型")
		return
	}
	//检查是否为结构体指针
	field := ptr.Elem()
	if field.Kind() != reflect.Struct {
		Logger.Error("sqlstr:[", sqlstr, "]的传入res类型不正确,传入的res的类型为:", ptr.Kind(), ",实际上应当为结构体指针")
		err = errors.New("res应当为结构体指针")
		return
	}
	result, err := dbPool.Query(sqlstr)
	if err != nil {
		Logger.Error("查询sql失败", err.Error())
		return
	}
	columns, err := result.Columns()
	if err != nil {
		Logger.Error("查询sql结果集的列的数量失败", err.Error())
		return
	}
	if err != nil {
		Logger.Error("查询sql结果集的列的类型失败", err.Error())
		return
	}
	scanArgs := make([]interface{}, len(columns))
	//vauleArgs := make([]string, len(columns))
	//
	//for i, _ := range scanArgs {
	//	var inter interface{}
	//	scanArgs[i] = &inter
	//}
loop:
	for i, dbField := range columns {
		for j := 0; j < field.NumField(); j++ {
			if field.Type().Field(j).Tag.Get("sqlstr") == dbField {
				scanArgs[i] = field.Field(j).Addr().Interface()
				continue loop
			}
		}
		var inter interface{}
		scanArgs[i] = &inter
	}

	for result.Next() {
		result.Scan(scanArgs...)
		//for _, arg := range scanArgs {
		//	fmt.Println((*arg.(*string)))
		//}
		break
	}
	return
}
