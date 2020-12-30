//pass_fail reports whether a grade is passing or failing

package main

import (
	"fmt"
	"log"
	"masuda.com/keyboard"
)

func main() {
	fmt.Print("Digite a nota: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	// Criando variavel dentro do escopo da func main, para que seja acessada no final do script
	var status string

	// If else comparando se a var grade é maior ou igual 60
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}

	// Output resultado final
	fmt.Println("You have: ", status)
}
