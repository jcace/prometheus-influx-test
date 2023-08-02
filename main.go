package main

import "fmt"

func main() {
	err := PostToPrometheus()

	if err != nil {
		fmt.Println("failure: " + err.Error())
	} else {
		fmt.Println("success!")
	}
}
