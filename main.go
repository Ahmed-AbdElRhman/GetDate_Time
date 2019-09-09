package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(GetDate())
	fmt.Println(GetTime())

}

///GetDate return Date (String)
func GetDate() string {
	return time.Now().Format("2006/01/02")
}

func GetTime() string {
	return time.Now().Format(time.Kitchen)
}
