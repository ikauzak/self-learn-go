package main

import (
	"fmt"
	"string"
)

func main() {
	// Criando var broken
	broken := "G# r#cks!"

	// Criando um objeto do string com método NewReplacer(antigo, novo)
	replacer := strings.NewReplacer("#", "o")

	// Executando a função Replace com o objeto replacer e passando a variavel broken como argumento
	fixed := replacer.Replace(broken)

	// Print
	fmt.Println(fixed)

	// Chamando a função Replace diretamente usando o objeto replacer e passando o argumento broken
	fmt.Println(replacer.Replace(broken))

	// Aqui a execução ocorre em uma linha apenas
	fmt.Println(strings.NewReplacer("#", "o").Replace(broken))
}
