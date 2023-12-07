package main

import (
	"flag"
	"fmt"

	csvsorter "github.com/GabrielMessiasdaRosa/desafio-go-imersao-fullcycle/internal/csv-sorter"
)

func main() {
	sortField := flag.String("sort", "", "Campo para ordenação")
	order := flag.String("order", "asc", "Ordem de ordenação (ascendente ou descendente)")
	flag.Parse()

	if *sortField == "" {
		fmt.Println("Por favor, forneça um campo para ordenação usando a flag --sort")
		return
	}

	if len(flag.Args()) != 2 {
		fmt.Println("Uso: go run main.go --sort=Campo_Ordenacao --order=asc|desc arquivo_entrada.csv arquivo_saida.csv")
		return
	}

	err := csvsorter.ExecuteSortCSV(flag.Args()[0], flag.Args()[1], *sortField, *order)
	if err != nil {
		fmt.Println("Erro ao ordenar o CSV:", err)
		return
	}

	fmt.Println("Arquivo CSV ordenado com sucesso.")
}
