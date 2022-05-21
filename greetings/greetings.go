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

// Hello retorna um mapa que associa cada uma das pessoas pessoa nomeada.
// com uma mensagem de saudação.
func Hellos(names []string) (map[string]string, error) {
	// Um mapa para associar nomes a mensagens.
	messages := make(map[string]string)
	// Percorre a fatia de nomes recebida, chamando
	// a função Hello para obter uma mensagem para cada nome.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// No mapa, associe a mensagem recuperada ao nome.
		messages[name] = message
	}
	return messages, nil
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
