package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/zohaibsoomro/go-quiz/models"
)

func ProblemPuller(fileName string) []models.Problem {
	problems := []models.Problem{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening file: ", err.Error())
		return nil
	}
	defer file.Close()
	pbs, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error while reading file: ", err.Error())
		return nil
	}
	for _, qAndA := range pbs {
		problem := parseProblem(qAndA)
		problems = append(problems, *problem)
	}
	return problems
}

func parseProblem(qAndA []string) *models.Problem {
	return &models.Problem{
		Question: qAndA[0],
		Answer:   qAndA[1],
	}
}
