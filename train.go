package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	var a2 = []rune(a)
	var sravn = []rune("")
	for i := len(a2) - 1; i >= 0; i-- {
		sravn = append(sravn, a2[i])
	}
	if string(sravn) == string(a2) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
