package webengine

import (
	"reflect"
)

func QueryRow(sqlstr string, res interface{}) (result *SqlResult, err error) {
	//检查是否为结构体
	value := reflect.ValueOf(res)
	if !mustBeStruct(value) {
		return
	}
	//执行sql
	row, err := dbPool.Query(sqlstr)
	defer row.Close()
	if err != nil {
		Logger.Error("Query失败:", err.Error())
		return
	}
	columns, err := row.Columns()
	if err != nil {
		Logger.Error("查询sql结果集的列的信息失败:", err.Error())
		return
	}
	scanArgs := make([]interface{}, len(columns))
	for i, _ := range scanArgs {
		var inter interface{}
		scanArgs[i] = &inter
	}

	result = new(SqlResult)

	//构造一个临时对象
	correspond := trivalStructField(columns, value)
	resultValue := reflect.New(reflect.TypeOf(res))
	bindInterfaceSlice(&scanArgs, resultValue, correspond)

	for row.Next() {
		result.RowCount++
		row.Scan(scanArgs...)
		result.ResultSlice = resultValue.Elem().Interface()
		break
	}
	return
}

func Query(sqlstr string, res interface{}) (result *SqlResult, err error) {
	//检查是否为结构体
	value := reflect.ValueOf(res)
	if !mustBeStruct(value) {
		return
	}
	//执行sql
	row, err := dbPool.Query(sqlstr)
	defer row.Close()
	if err != nil {
		Logger.Error("Query失败:", err.Error())
		return
	}
	columns, err := row.Columns()
	if err != nil {
		Logger.Error("查询sql结果集的列的信息失败:", err.Error())
		return
	}
	scanArgs := make([]interface{}, len(columns))
	for i, _ := range scanArgs {
		var inter interface{}
		scanArgs[i] = &inter
	}

	result = new(SqlResult)

	//构造一个临时对象
	correspond := trivalStructField(columns, value)
	resultValue := reflect.New(reflect.TypeOf(res))
	bindInterfaceSlice(&scanArgs, resultValue, correspond)
	tmpSliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(res)), 0, 0)

	for row.Next() {
		row.Scan(scanArgs...)
		tmpSliceValue = reflect.Append(tmpSliceValue, resultValue.Elem())
		result.RowCount++
	}
	result.ResultSlice = tmpSliceValue.Interface()
	return
}
