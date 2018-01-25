package main

import "fmt"
import bf "bufio"
import "os"
import "strconv"

func main() {
	var input = bf.NewScanner(os.Stdin)

	fmt.Println("Input the number that will be guessed:")
	input.Scan()

	var input1, _ = strconv.Atoi(input.Text())
	fmt.Print("\033[H\033[2J")
	fmt.Println("make a guess:")
	input.Scan()
	var input21, _ = strconv.Atoi(input.Text())

	for {
		if input21 == input1 {
			fmt.Println("You win")
			break
		}
		if input21 > input1 {
			fmt.Println("Too high, try agian")
			input.Scan()
			input21, _ = strconv.Atoi(input.Text())
		}
		if input21 < input1 {
			fmt.Println("Too low, try again")
			input.Scan()
			input21, _ = strconv.Atoi(input.Text())
		}
	}
}
