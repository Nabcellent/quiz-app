package quiz_app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	q string
	a string
}

func parseProblems(lines [][]string) []Problem {
	//	Go over lines and parse them, with Problem struct
}

func problemPuller(fileName string) ([]Problem, error) {
	//	Read all problems from csv.

	//	1.	Open file
	if file, err := os.Open(fileName); err != nil {
		//	2.	Create new reader
		csvR := csv.NewReader(file)

		//	3.	Read file
		if cLines, err := csvR.ReadAll(); err != nil {

			//	4.	Call problem parser func
			return parseProblems(cLines), nil
		} else {
			return nil, fmt.Errorf("error while reading data in csv format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error while opening %s file; %s", fileName, err.Error())
	}
}

func main() {
	//	1.	Input name of file.
	fileName := flag.String("f", "quiz.csv", "Path to csv file.")

	//	2.	Set duration of timer.
	timer := flag.Int("t", 30, "Timer for the quiz.")
	flag.Parse()

	//	3.	Pull problems from file (call problem puller function).
	problems, err := problemPuller(*fileName)

	//	4.	Handle error.
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	//	5.	Create a variable to count correct answers.
	correctAnswer := 0

	//	6.	Initialize timer using duration set in step 2.
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

	//	7.	Loop through the problems, print the questions and accept answers.

	//	8.	Calculate and print out result.

}
