package main

import (
	"fmt"
)

func printNumber(numChannel, letterChannel *chan bool) {
	i := 1
	for {
		select {
		case <-(*numChannel):
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			(*letterChannel) <- true
			break
		default:
			break
		}
	}
}

func printLetter(numChannel, letterChannel *chan bool) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	i := 0
	//(*numChannel) <- true
	for {
		select {
		case <-(*letterChannel):
			if i >= len(str) {
				return
			}
			fmt.Print(str[i : i+1])
			i++
			if i >= len(str) {
				i = 0
			}
			fmt.Print(str[i : i+1])
			i++
			(*numChannel) <- true
			break
		default:
			break
		}
	}
}

func main() {
	letter, number := make(chan bool), make(chan bool)
	go printNumber(&number, &letter)
	
	number <- true // ?????
	printLetter(&number, &letter)
	fmt.Println()
}
