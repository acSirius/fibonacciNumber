package main
import "fmt"

func main () {
	const fibonacciLength int = 80 // How long do you want the sequence to be ?
	var fibonacciNumbers [fibonacciLength]int // Array of fibonacci numbers with the desired length

	fibonacciNumbers[0], fibonacciNumbers[1] = 0, 1 // First two values of fibonacci, this need to be hard coded to be able to find the rest

	for i := 0; i < fibonacciLength;i++ {
		if i < 2{ fmt.Println(fibonacciNumbers[i]) } else {
			fibonacciNumbers[i] = fibonacciNumbers[i - 2] + fibonacciNumbers[i - 1]
			fmt.Println(fibonacciNumbers[i])
		}
	}

}
