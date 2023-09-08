package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	
	chislo := 1

	for chislo <= 100{

		fmt.Println(fizzbuzz.FizzBuzz(chislo))

		chislo++

	}	
}
