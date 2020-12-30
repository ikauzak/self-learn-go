package main

import (
	"fmt"
)

/*
   Pointers (&) sao utilizados para apontar o endereço de uma var
   na memoria do programa.

   O simbolo "*" sao utilizados para criar variaveis que seguram
   valores de pointers.
*/
func main() {

	var myInt int              //Criando var do tipo int
	var myIntPointer *int      //Criando uma var que vai receber um pointer do tipo int
	myIntPointer = &myInt      //Atribuindo o endereco da variavel myInt em myIntPointer
	myInt = 42                 //Atribuindo um valor a variavel myInt
	fmt.Println(myIntPointer)  //Imprimindo o valor da variavel myIntPointer(endereco de memoria)
	fmt.Println(*myIntPointer) //Imprimindo o valor apontado pelo endereco da memoria
}
