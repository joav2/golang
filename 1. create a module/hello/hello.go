package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Configura as propriedades do Logger predefinido, incluindo
	// o prefixo da entrada de log e um sinalizador para desabilitar a impressão
	// da hora, arquivo de origem e número da linha.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Uma fatia de nomes.
	names := []string{"Lucieny", "João Vitor", "Darrin"}

	// Solicita uma mensagem de saudação.
	message, err := greetings.Hellos(names)
	// Se um erro foi retornado, imprima-o no console e
	// saia do programa.
	if err != nil {
		log.Fatal(err)
	}

	// Se nenhum erro foi retornado, imprima a mensagem retornada
	// no console.
	fmt.Println(message)
}
