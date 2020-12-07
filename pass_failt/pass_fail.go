//pass_fail reports whether a grade is passing or failing

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Digite a nota: ")

	// Guardar input do usuário
	reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// Guardar a string digitada até o "enter" (\n)
	input, err := reader.ReadString('\n')

	// Verificar se há conteúdo no err

	if err != nil {
		log.Fatal(err)
	}

	// Remover o \n enter do usuário
	input = strings.TrimSpace(input)

	// Converter a variavel input (antiga string) em valor float 64
	grade, err := strconv.ParseFloat(input, 64)

	fmt.Println("Your score was: ", input)

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
