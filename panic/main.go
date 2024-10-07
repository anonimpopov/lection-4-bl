package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("handled: %s", err)
		}
	}()
	go func() {
		panic("panic")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("end")
}
