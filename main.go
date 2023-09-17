package main

import (
	"fmt"
	"time"

	"github.com/zohaibsoomro/go-quiz/utils"
)

func main() {
	problems := utils.ProblemPuller("quiz.csv")
	qNo := 0
	correctAnswers := 0
	timer := time.NewTimer(time.Second * 30)

	fmt.Println("Quiz starting...")
quizLoop:
	for {
		pb := problems[qNo]
		answers := make(chan string)
		fmt.Printf("Q#%d: %s = ", qNo+1, pb.Question)
		go func(c chan string) {
			var answer string
			fmt.Scan(&answer)
			answers <- answer
		}(answers)

		select {
		case answer := <-answers:
			if answer == pb.Answer {
				correctAnswers++
				fmt.Println("✅")
			} else {
				fmt.Println("❌")
			}
		case <-timer.C:
			fmt.Println("Time out ⏳..")
			break quizLoop
		}

		if qNo == len(problems)-1 {
			break quizLoop
		}
		qNo++
	}
	fmt.Println("Quiz finished.")
	fmt.Printf("Your score is %d.\n", correctAnswers)
}
