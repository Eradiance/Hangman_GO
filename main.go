package main

import (
	l "Hangman/words"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Виселица запущена")
	words := l.New()
	for {
		paramMenu := getMenu()
		if paramMenu != "1" {
			break
		}
		startNewGame(words)
	}
}

func getMenu() string {
	fmt.Println("1. Начать игру")
	fmt.Println("2. Выход")
	return promptData("Введите выбор")
}

func promptData(param string) string {
	var data string
	fmt.Print(param + " ")
	fmt.Scan(&data)
	return data
}

func startNewGame(words *l.ListOfWords) {
	word := strings.ToLower(words.GetRandomWord())
	var char string
	var healthPint int = 0

	secretWord := make([]rune, len(word))
	for i := range word {

		secretWord[i] = rune('_')
	}

	for {
		healthFlag := true
		fmt.Print("Введите букву: ")
		fmt.Scan(&char)

		for i, val := range word {
			if string(unicode.ToLower(val)) == char {
				secretWord[i] = val
				healthFlag = false
			}
		}

		if healthFlag {
			healthPint += 1
		}
		fmt.Println(hangmanArts[healthPint])
		fmt.Println(string(secretWord))

		if healthPint == 7 {
			fmt.Println("Вы проиграли! Загаданное слово: ", word)
			break
		}
	}

	fmt.Println()
}
