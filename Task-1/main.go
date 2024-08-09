package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculate average grade function
func calculateAverageGrade(grades []float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var studentName string
	fmt.Print("Enter Your Name (or 'q' to quit): ")
	reader := bufio.NewReader(os.Stdin)
	studentName, _ = reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	if strings.ToLower(studentName) == "q" {
		fmt.Println("Exiting the program.")
		return
	}

	var numberOfSubjects int

	for {
		fmt.Print("Enter the number of subjects (or 'q' to quit): ")
		var userInput string
		reader := bufio.NewReader(os.Stdin)
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "q" {
			fmt.Println("Exiting the program.")
			return
		}
		
		numberOfSubjects, err := strconv.Atoi(userInput)
		if err != nil || numberOfSubjects <= 0 {
			fmt.Println("Invalid input. Please enter a positive integer.")
			continue
		}
		break
	}


	subjectGrades := make(map[string]float64)

	for i := 0; i < numberOfSubjects; i++ {
		var subjectName string
		var grade float64
		fmt.Printf("Enter the name of subject %d (or 'q' to quit): ", i+1)
		reader := bufio.NewReader(os.Stdin)
		subjectName, _ = reader.ReadString('\n')
		subjectName = strings.TrimSpace(subjectName)

		if strings.ToLower(subjectName) == "q" {
			fmt.Println("Exiting the program.")
			return
		}

		for {
			fmt.Printf("Enter the grade for %s (or 'q' to quit): ", subjectName)
			var gradeInput string
			
			reader := bufio.NewReader(os.Stdin)
			gradeInput, _ = reader.ReadString('\n')
			gradeInput = strings.TrimSpace(gradeInput)

			if strings.ToLower(gradeInput) == "q" {
				fmt.Println("Exiting the program!")
				return
			}

			grade, err := strconv.ParseFloat(gradeInput, 64)
			if err != nil || grade < 0 || grade > 100 {
				fmt.Println("Invalid input. Please enter a grade between 0 and 100.")
				continue
			}
			break
		}

		subjectGrades[subjectName] = grade
	}


	grades := make([]float64, 0, numberOfSubjects)
	for _, grade := range subjectGrades {
		grades = append(grades, grade)
	}

	// Calculating the average grade
	averageGrade := calculateAverageGrade(grades)

	
	fmt.Println("\n~~~~~ Student Grade Report ~~~~~")
	fmt.Printf("Student Name: %s\n", studentName)
	for subject, grade := range subjectGrades {
		fmt.Printf("Subject: %s, Grade: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}
