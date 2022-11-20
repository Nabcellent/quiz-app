package main

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
	r := make([]Problem, len(lines))

	for i := 0; i < len(lines); i++ {
		r[i] = Problem{q: lines[i][0], a: lines[i][1]}
	}

	return r
}

func problemPuller(fileName string) ([]Problem, error) {
	//	Read all problems from csv.

	//	1.	Open file
	if file, err := os.Open(fileName); err == nil {
		//	2.	Create new reader
		csvR := csv.NewReader(file)

		//	3.	Read file
		if cLines, err := csvR.ReadAll(); err == nil {

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
	correctAns := 0

	//	6.	Initialize timer using duration set in step 2.
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

	//	7.	Loop through the problems, print the questions and accept answers.
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problems %d: %s = ", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)

			ansC <- answer
		}()

		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}

			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}

	//	8.	Calculate and print out result.
	fmt.Printf("You result is %d out of %d\n", correctAns, len(problems))
	fmt.Printf("Press enter to exit.")

	<-ansC
}

func exit(msg string) {
	fmt.Println(msg)

	os.Exit(1)
}
