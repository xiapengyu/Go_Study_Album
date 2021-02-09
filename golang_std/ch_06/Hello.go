package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Hello(""))
	Say("english")
}

func Hello(message string) string {
	if message == "" {
		message = "Hello"
	}
	return message + " World!"
}

func Say(language string)  {
	switch language{
	case "french":
		fmt.Println("I come from french")
	case "english":
		fmt.Println("I come from english")
	default:
		fmt.Println("I come from China")
	}
}
