package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Ola mundo!"
	canal <- "Programando em Go"

	msg := <-canal
	msg2 := <-canal
	msg3 := <-canal
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println(msg)
}
