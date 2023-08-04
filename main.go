package main

import (
	"context"
	"fmt"
)

func main() {
	err := PostToPrometheus(context.TODO())

	if err != nil {
		fmt.Println("failure: " + err.Error())
	} else {
		fmt.Println("success!")
	}
}
