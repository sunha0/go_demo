package main

import (
	"fmt"
	"reflect"
)

func reflectType(arg interface{}) {

	t := reflect.TypeOf(arg)

	fmt.Println(t, t.Name(), t.Kind())

}

func reflectValue(arg interface{}) {

	v := reflect.ValueOf(arg)

	fmt.Printf("%T %v \n", v, v)

	k := v.Kind()

	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%T %v \n", ret, ret)
	case reflect.String:
		ret := string(v.String())

		fmt.Printf("%T %v \n", ret, ret)
	}

}

func reflectSetValue(arg interface{}) {

	v := reflect.ValueOf(arg)

	fmt.Printf("%T %v \n", v, v)

	k := v.Elem().Kind()

	switch k {
	case reflect.Float32:
		v.Elem().SetFloat(3.21)

	case reflect.String:

		v.Elem().SetString("嘿嘿")

	}

}

type Dog struct{}

func main() {

	var a float32 = 1.643

	reflectType(a)
	var b int32 = 64
	reflectType(b)

	var c = map[string]string{"name": "xiaoMing", "gender": "male"}

	reflectType(c)

	var d = Dog{}
	reflectType(d)

	fmt.Println("----------------------")
	reflectValue(a)

	reflectValue("哈哈")

	fmt.Println("----------------------")

	reflectSetValue(&a)
	fmt.Println(a)
	s := "哈哈"
	reflectSetValue(&s)
	fmt.Println(s)
}
