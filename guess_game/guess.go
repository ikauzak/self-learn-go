package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Criando seed para gerar número aletório
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	// Criando um numero aleatorio entre 1 e 100
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	// Linha para mostrar valor do numero criado fmt.Println(target)

	// COletar input do usuario
	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		// Removendo o \n do usuario
		input = strings.TrimSpace(input)

		// Convertendo o valor string para int
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	// Se a var success for diferente de true
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
