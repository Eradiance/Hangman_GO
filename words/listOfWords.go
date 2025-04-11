package words

import "math/rand"

type ListOfWords struct {
	words []string
}

func New() *ListOfWords {
	return &ListOfWords{
		words: []string{
			"Одуванчик",
			"Ромашка",
			"Машина",
			"Компьютер",
			"Телефон",
			"Работа",
		},
	}
}

func (listOfWords *ListOfWords) GetRandomWord() string {
	return listOfWords.words[rand.Intn(len(listOfWords.words))]
}
