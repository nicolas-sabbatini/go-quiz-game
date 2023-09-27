package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	correct  string
}

func (q *Question) ask() {
	fmt.Println(q.question)
}

func (q *Question) isCorrect(answer string) bool {
	return q.correct == answer
}

type Quiz struct {
	questions []Question
	correct   int
	time      time.Duration
	finished  bool
}

func (q *Quiz) askQuiz(done chan bool) {
	for number, question := range q.questions {
		fmt.Println("\nQuestion #", number+1)
		question.ask()
		if question.isCorrect(awaitAnswer()) {
			q.correct++
			fmt.Println("🎉 ¡Correct! 🎉")
		} else {
			fmt.Println("❌ Incorrect ❌")
		}
		fmt.Println("You have", q.correct, "correct answers")
	}
	q.finished = true
	done <- true
}

// func (q *Quiz) startTimer(done chan bool) {
// 	time.Sleep(q.time)
// 	done <- true
//}

func (q *Quiz) startTimer() <-chan time.Time {
	return time.NewTimer(q.time).C
}

func (q *Quiz) finishGame() {
	if !q.finished {
		fmt.Println("🕑 Time is up! 🕑 Better luck next time!")
	}
	q.finished = true
	fmt.Println("\nYou have", q.correct, "correct answers out of", len(q.questions), "questions")
	if q.correct == len(q.questions) {
		fmt.Println("🎉🎉🎉 Congratulations you are a GENIUS 🎉🎉🎉")
	} else if q.correct > len(q.questions)/2 {
		fmt.Println("👍 You are amazing 👍")
	} else {
		fmt.Println("😢 You lose, try again 😢")
	}
}

func isFatalError(message string, err error) {
	if err != nil {
		log.Fatal(message, " : ", err)
	}
}

func parseArgs() string {
	path := flag.String("path", "", "Path to the CSV file")
	flag.Parse()
	if *path == "" {
		log.Fatal("Please provide a file name")
	}
	return *path
}

func loadQuiz(path string) Quiz {
	quizCsv, err := os.Open(path)
	isFatalError("Error opening file", err)
	defer quizCsv.Close()
	csvRows, err := csv.NewReader(quizCsv).ReadAll()
	isFatalError("Error reading CSV file", err)
	var questions []Question
	duration, err := time.ParseDuration(csvRows[0][1])
	isFatalError("Error parsing duration", err)
	for _, row := range csvRows[1:] {
		questions = append(questions, Question{
			question: row[0],
			correct:  row[1],
		})
	}
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
	quiz := Quiz{
		questions: questions,
		correct:   0,
		time:      duration,
		finished:  false,
	}
	return quiz
}

func awaitAnswer() string {
	in := bufio.NewReader(os.Stdin)
	answer, err := in.ReadString('\n')
	isFatalError("Error reading answer", err)
	answer = strings.TrimSuffix(answer, "\n")
	if answer == "" {
		return awaitAnswer()
	}
	return answer
}

func main() {
	quiz := loadQuiz(parseArgs())
	done := make(chan bool)
	go quiz.askQuiz(done)
	// go quiz.startTimer()
	timer := quiz.startTimer()
	select {
	case <-timer:
	case <-done:
	}
	quiz.finishGame()
}
