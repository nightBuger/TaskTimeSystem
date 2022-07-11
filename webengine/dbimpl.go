package webengine

import (
	"errors"
	"fmt"
	"reflect"
)

func Query(sql string, res interface{}) (err error) {
	//检查是否为指针
	ptr := reflect.ValueOf(res)
	if ptr.Kind() != reflect.Ptr {
		Logger.Error("sql:[", sql, "]的传入res类型不正确,传入的res的类型为:", ptr.Kind(), ",实际上应当为指针类型")
		err = errors.New("res应当为指针类型")
		return
	}
	//检查是否为结构体指针
	field := ptr.Elem()
	if field.Kind() != reflect.Struct {
		Logger.Error("sql:[", sql, "]的传入res类型不正确,传入的res的类型为:", ptr.Kind(), ",实际上应当为结构体指针")
		err = errors.New("res应当为结构体指针")
		return
	}
	result, err := dbPool.Query(sql)
	if err != nil {
		Logger.Error("查询sql失败", err.Error())
		return
	}
	columns, err := result.Columns()
	if err != nil {
		Logger.Error("查询sql结果集的列的数量失败", err.Error())
		return
	}
	columnTypes, err := result.ColumnTypes()
	for _, columnType := range columnTypes {
		fmt.Println(columnType.Name())
		fmt.Println(columnType.ScanType())
		fmt.Println(columnType.DatabaseTypeName())
		fmt.Println(columnType.Length())
	}
	if err != nil {
		Logger.Error("查询sql结果集的列的类型失败", err.Error())
		return
	}
	scanArgs := make([]interface{}, len(columns))
	for result.Next() {
		result.Scan(scanArgs...)
	}
	return
}
