
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	ID            int               `json:"id"`
	Difficulty    string            `json:"difficulty"`
	Question      string            `json:"question"`
	Options       map[string]string `json:"options"`
	CorrectOption string            `json:"correct_option"`
}

type AptitudeData struct {
	AptitudeQuestions []Question `json:"AptitudeQuestions"`
}

// Function to handle taking a quiz with the provided questions
func takeQuiz(questions []Question, size int) int {
	score := 0
	fmt.Println("Starting the test...")

	for i := 0; i < size; i++ {
		question := questions[i]
		fmt.Println("Question:", question.Question)
		fmt.Println("Options:")
		for key, value := range question.Options {
			fmt.Printf("%s: %s\n", key, value)
		}

		var answer string
		for {
			fmt.Println("Please select an option (A, B, C, D):")
			fmt.Scan(&answer)
			if strings.ContainsAny(strings.ToUpper(answer), "ABCD") && len(answer) == 1 {
				break
			} else {
				fmt.Println("Invalid option. Please try again.")
			}
		}
		answer = strings.ToUpper(answer)

		if answer == question.CorrectOption {
			fmt.Println("Correct Answer!")
			score++

		} else {
			fmt.Println("Wrong Answer!")
			fmt.Printf("The correct answer was: %s: %s\n",
				question.CorrectOption,
				question.Options[question.CorrectOption])
		}
	}

	fmt.Printf("Test Ended successfully with a score of: %d\n", score)
	return score
}

// Function to get questions by difficulty
func getQuestionsByDifficulty(allQuestions []Question, difficulty string) []Question {
	var filteredQuestions []Question
	for _, question := range allQuestions {
		if question.Difficulty == difficulty {
			filteredQuestions = append(filteredQuestions, question)
		}
	}
	return filteredQuestions
}

// Function to get user-selected number of questions
func getNumberOfQuestions(maxQuestions int) int {
	var size int
	for {
		fmt.Println("Enter the number of questions you want to solve:")
		_, err := fmt.Scanln(&size)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if size > maxQuestions || size <= 0 {
			fmt.Printf("Invalid number of questions selected. Please enter a number between 1 and %d\n", maxQuestions)
			continue
		}
		break
	}
	return size
}

// Function to get random questions
func getRandomQuestions(allQuestions []Question, count int) []Question {
	length := len(allQuestions)
	if count > length {
		count = length
	}

	// Create a copy of the questions to shuffle
	shuffled := make([]Question, length)
	copy(shuffled, allQuestions)

	// Fisher-Yates shuffle
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Random Question ID:", shuffled[i].ID)
	}

	return shuffled[:count]
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Open and read the JSON file
	jsonFile, err := os.Open("Aptitude.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer jsonFile.Close()

	readContent, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data AptitudeData
	err = json.Unmarshal(readContent, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	allQuestions := data.AptitudeQuestions

	for i := 0; i < 10; i++ {
		fmt.Println("Random Question ID:", allQuestions[i].ID)
	}

	// Display menu
	fmt.Println("Select an option to start the test")
	fmt.Println("1] Time limit Aptitude Test")
	fmt.Println("2] Solve Questions Randomly")
	fmt.Println("3] Solve Easy Questions")
	fmt.Println("4] Solve Medium Questions")
	fmt.Println("5] Solve Hard Questions")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:

		fmt.Println("Time limit Aptitude Test")
		fmt.Println("Enter the time limit in minutes")
		var timeLimit int
		fmt.Scanln(&timeLimit)
		fmt.Printf("You have selected a time limit of %d minutes\n", timeLimit)

		// Shuffle all questions
		shuffledQuestions := getRandomQuestions(allQuestions, len(allQuestions))

		score := 0
		fmt.Println("Starting the test...")

		// Set up the timer
		startTime := time.Now()
		endTime := startTime.Add(time.Duration(timeLimit) * time.Minute)

		questionIndex := 0
		totalQuestions := len(shuffledQuestions)

		for questionIndex < totalQuestions {
			// Check if time is up before each question
			if time.Now().After(endTime) {
				fmt.Println("\n⏰ Time's up! Exiting...")
				break
			}

			// Display time remaining
			remaining := endTime.Sub(time.Now())
			mins := int(remaining.Minutes())
			secs := int(remaining.Seconds()) % 60
			fmt.Printf("\nTime remaining: %d min %d sec\n", mins, secs)

			question := shuffledQuestions[questionIndex]
			fmt.Printf("Question %d/%d: %s\n", questionIndex+1, totalQuestions, question.Question)
			fmt.Println("Options:")
			for key, value := range question.Options {
				fmt.Printf("%s: %s\n", key, value)
			}

			// Create a channel to receive user input
			answerChan := make(chan string)

			// Goroutine to get user input
			go func() {
				var answer string
				fmt.Println("Please select an option (A, B, C, D):")
				fmt.Scan(&answer)
				answerChan <- answer
			}()

			// Wait for either user input or timer expiration
			select {
			case answer := <-answerChan:
				answer = strings.ToUpper(answer)
				if strings.ContainsAny(answer, "ABCD") && len(answer) == 1 {
					if answer == question.CorrectOption {
						fmt.Println("Correct Answer!")
						score++

					} else {
						fmt.Println("Wrong Answer!")
						fmt.Printf("The correct answer was: %s: %s\n",
							question.CorrectOption,
							question.Options[question.CorrectOption])
					}
					questionIndex++
				} else {
					fmt.Println("Invalid option. Please try again.")
					// Stay on the same question
				}

			case <-time.After(endTime.Sub(time.Now())):
				fmt.Println("\n⏰ Time's up! Exiting...")
				goto timerEnd // Break out of both loops
			}
		}

	timerEnd:
		fmt.Printf("\nTest Ended with a score of: %d/%d\n", score, questionIndex)
		time.Sleep(2 * time.Second)

	case 2:
		fmt.Println("Solve Questions Randomly")
		length := len(allQuestions)
		fmt.Println("Available Questions:", length)

		size := getNumberOfQuestions(length)
		fmt.Printf("You have selected to solve %d questions randomly\n", size)

		randomQuestions := getRandomQuestions(allQuestions, size)
		takeQuiz(randomQuestions, size)

	case 3:
		fmt.Println("Solve Easy Questions")
		easyQuestions := getQuestionsByDifficulty(allQuestions, "Easy")
		easySize := len(easyQuestions)
		fmt.Println("Available Easy Questions:", easySize)

		if easySize == 0 {
			fmt.Println("No Easy questions available")
			return
		}

		size := getNumberOfQuestions(easySize)
		fmt.Printf("You have selected to solve %d Easy questions\n", size)

		takeQuiz(easyQuestions, size)

	case 4:
		fmt.Println("Solve Medium Questions")
		mediumQuestions := getQuestionsByDifficulty(allQuestions, "Medium")
		mediumSize := len(mediumQuestions)
		fmt.Println("Available Medium Questions:", mediumSize)

		if mediumSize == 0 {
			fmt.Println("No Medium questions available")
			return
		}

		size := getNumberOfQuestions(mediumSize)
		fmt.Printf("You have selected to solve %d Medium questions\n", size)

		takeQuiz(mediumQuestions, size)

	case 5:
		fmt.Println("Solve Hard Questions")
		hardQuestions := getQuestionsByDifficulty(allQuestions, "Hard")
		hardSize := len(hardQuestions)
		fmt.Println("Available Hard Questions:", hardSize)

		if hardSize == 0 {
			fmt.Println("No Hard questions available")
			return
		}

		size := getNumberOfQuestions(hardSize)
		fmt.Printf("You have selected to solve %d Hard questions\n", size)

		takeQuiz(hardQuestions, size)

	default:
		fmt.Println("Invalid option selected")
	}
}
