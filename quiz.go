package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}

type Question struct {
	question string
	answer   string
}

func readCsv(filename string) []Question {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records[0])
	// Допишите код здесь
	return nil
}

func main() {
	total := 0

	questions := []Question{{"1 + 1", "2"}, {"2 + 2", "4"}}
	// Пройтись циклом. Вывести вопрос, предложить пользователю ввести ответ.
	// Если ответ правильный, увеличить total.
	// for
}
