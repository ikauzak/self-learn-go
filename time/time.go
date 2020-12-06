package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando uma var com nome now do time Time usando a lib time com o valor usando a função Now dentro da lib time
	var now time.Time = time.Now()
	// Print do valor "now" de forma crua
	fmt.Println(now)
	//	var year int = now.Year() -> Chamando a função Year com o objeto criado now
	fmt.Println(now.Year())
	// Chamando a função Hour com o objeto "now" criado na linha 10
	fmt.Println(now.Hour())
	// Chamando a função Clock com o objeto "now" criado na linha 10
	fmt.Println(now.Clock())

	// Criando uma nova variavel chamada date do tipo string, usando a função Format para formatar padrão de data
	// Por padrão no GO, o valor para formatar data é Mon Jan 2 15:04:05 -0700 MST 2006
	// Onde cada Mon representa o dia da semana, Jan representa o mês e assim por diante
	var date string = now.Format("15:04:05 2/Jan/2006")
	fmt.Println(date)

	// Chamando a função "Format" diretamente através do Println para printar houra:minuto:segundo dia/mês/ano
	fmt.Println(now.Format("15:04:05 - - - 2/Dec/2006"))
}
