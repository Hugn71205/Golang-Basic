package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func randomlc() string {
	randomloichao := []string{
		"Xin chao ngay moi %v",
		"Hom nay ban the nao %v",
	}
	return randomloichao[rand.IntN(len(randomloichao))]
}
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Ban da de trong ten!")
	}
	message := fmt.Sprintf(randomlc(), name)
	return message, nil
}
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
