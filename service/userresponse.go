package service

import (
	"fmt"
	"reflect"
)

type UserInfoGetRes struct {
	Id       string      `json:"id" sql:"userid"`
	Name     string      `json:"name" sql:"username"`
	Level    int         `json:"level" sql:"secretlevel"`
	Email    string      `json:"email" sql:"email"`
	Phonenum string      `json:"phonenum" sql:"lxdh"`
	userinfo UserInfoGet `json:"userinfo" sql:"userinfo"`
	roleid   []string    `json:"roleid" sql:"roleid"`
}

func getReflectInfo(any interface{}) {
	var value reflect.Value
	value = reflect.ValueOf(any)

	fmt.Println(value.Type(), value.Kind() == reflect.Ptr)
	value = value.Elem()
	fmt.Println(value.Type(), value.Kind() == reflect.Struct)
	for i := 0; i < value.NumField(); i++ {
		fieldType := value.Type().Field(i)
		field := value.Field(i)
		fmt.Println(fieldType.Name, ":", fieldType.Tag.Get("sql"), field.Kind())
		switch field.Kind() {
		case reflect.String:
			field.SetString("啊哈哈哈")
		case reflect.Int:
			field.SetInt(123456)
		case reflect.Struct:
			fmt.Println("遇到了struct字段 下次再说咯")
		case reflect.Slice:
			fmt.Println("遇上了数组,看看怎么处理", field.Type().Elem())
		default:
			fmt.Println("不认识的类型")
		}
	}
}

//func init() {
//	var v UserInfoGetRes
//	getReflectInfo(&v)
//	fmt.Println("test v.Id", v.Id)
//	fmt.Println("test v.Name", v.Name)
//	fmt.Println("test v.Level", v.Level)
//	fmt.Println("test v.Email ", v.Email)
//	fmt.Println("test v.Phonenum ", v.Phonenum)
//}
