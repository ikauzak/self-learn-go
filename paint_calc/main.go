package main

import (
	"fmt"
	"log"
)

/* Na func principal, chamamos a func paintWall passando dois argumentos
   (4.2 e 3.0), pois a funcao paintWall espera receber 2 argumentos.
   Os valores retornados desta chamada da funcao serao armazenados em 2
   variaveis (amount e err) pois a func paintWall tambem retorna 2 valores
   (explicacao mais abaixo)
*/

func main() {
	var amount, total float64
	amount, err := paintWall(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	total += amount
	//	amount, err = paintWall(5.7, 3.5)
	//	fmt.Println(err)
	//	fmt.Printf("%0.2f liters needed\n", amount)
	//	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}

/* Funcao paintWall recebe 2 argumentos, width e heigth do tipo float64
   e retorna 2 valores, 1 do tipo float64 e outro do tipo error

   Nas condicoes if, se a condicao for verdadeira (width < 0 ou height < 0)
   entao o primeiro valor retornado sera 0 (tipo float64) e o segundo uma string
   (mensagem de erro)

   Caso as condioes sejam falsas, entao o valor retornado da funcao sera o resultado
   area e nulo (error nil)

*/

func paintWall(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a heigth of %0.2f is invalid", height)
	}
	area := width * height / 10.0

	return area, nil
}
