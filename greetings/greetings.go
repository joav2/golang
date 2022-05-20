package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello retorna uma saudação para a pessoa nomeada.
func Hello(name string) (string, error) {
	// Se nenhum nome foi fornecido, retorna um erro com uma mensagem.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Se um nome foi recebido, retorne um valor que incorpore o nome
	// em uma mensagem de saudação.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init define valores iniciais para variáveis usadas na função.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat retorna uma de um conjunto de mensagens de saudação.
// mensagem retornada é selecionada aleatoriamente.
func randomFormat() string {
	// Uma fatia de formatos de mesagem.
	formats := []string{
		"Hi, %v. WelCome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Retorna um formato de mensagem selecionado aleatoriamente especificando
	// um índice aleatório para a fatia de formatos.
	return formats[rand.Intn(len(formats))]
}
