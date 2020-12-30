package main

import (
	"fmt"
)

/* Tambem é possivel passar pointers em funcoes
   Desde que a funcao tenha argumentos que sao do tipo
   pointers (*int, *string, *float64) e o retorno seja
   com o simbolo "*"
   O argumento deve ser passado com o simbolo "&" na chamada
   da funcao, conforme linha 16.
*/
func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
