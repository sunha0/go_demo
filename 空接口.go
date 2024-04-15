package main

import (

	//"os"
	//"github.com/gin-gonic/gin"

	"fmt"
)

var myfunction = func(nums ...interface{}) {
	fmt.Print(nums)
	fmt.Printf("%T", nums)
}

func main() {
	// maps := make(map[string]interface{})

	// maps["a"] = 1
	// maps["b"] = "b"
	// maps["c"] = false

	maps := map[string]interface{}{
		"number": 1,
		"string": "c",
		"bool":   false,
	}

	myfunction(1, "Mystring", []string{"1", "2"}, maps)

}
