//pass_fail reports whether a grade is passing or failing

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Digite a nota: ")
	reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your score was: ", input)

	if input >= 60 {
		status := "Passed"
	} else {
		status := "Failed"
	}

	fmt.Println("You have: ", status)
}
