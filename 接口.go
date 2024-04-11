package main

import (

	//"os"
	//"github.com/gin-gonic/gin"

	"fmt"
)


// 定义接口
type Cloud interface {
	CreateECS()
}

// 定义结构体
type Aws struct {
	region string
}

// 定义结构体方法
func (aws *Aws) CreateECS() {
	fmt.Println(aws.region)
}

// 定义结构体工厂函数
func AwsFactory(region string) *Aws {

	return &Aws{
		region,
	}
}

// 第一种写法
func InitcloudType1(region string) (clt Cloud) {
	// 将结构体地址赋值给接口cloud,相当于第二种写法
	clt = AwsFactory(region)
	return

}

// 第二种写法
// 传参Cloud 接口对象
func InitcloudType2(clt Cloud) {

	clt.CreateECS()

}

func main() {

	// StartServer()

	// 第一种写法
	aws := InitcloudType1("法兰克福1")
	aws.CreateECS()

	// 第二种写法、先生成结构体对象
	aws1 := AwsFactory("法兰克福2")

	InitcloudType2(aws1)

}
