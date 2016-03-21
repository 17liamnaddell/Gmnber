package main

import "fmt"
import bf "bufio"
import "os"
import "strconv"

type SVC int

func main() {
	fmt.Println("Loaded")
	var input = bf.NewScanner(os.Stdin)
	var input2 = bf.NewScanner(os.Stdin)

	fmt.Println("Input the number that will be guessed:")
	input.Scan()

	var input1, err = strconv.Atoi(input.Text())
	qwerty := 1
	for qwerty < 40 {
		fmt.Println("\n")
		qwerty++
	}
	fmt.Println("\n make a guess:")
	input2.Scan()
	var input21, err1 = strconv.Atoi(input2.Text())

	for {
		if input21 == input1 {
			fmt.Println("You win")
			break
		}
		if input21 > input1 {
			fmt.Println("Too high, try agian")
			input2.Scan()
			input21, err1 = strconv.Atoi(input2.Text())
		}
		if input21 < input1 {
			fmt.Println("Too low, try again")
			input2.Scan()
			input21, err1 = strconv.Atoi(input2.Text())
		}
	}
	fmt.Println(err)
	fmt.Print(err1)
}
