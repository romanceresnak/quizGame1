package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct{
	question string
	answer string
}

func main(){
	csvFile := flag.String("csv","problem.csv","Parsong")
	timeLimit := flag.Int("limit",2,"the time limit")

	//set all the values
	flag.Parse()

	file,err := os.Open(*csvFile)

	if err != nil{
		exit("Error during opening file")
	}

	r := csv.NewReader(file)

	lines,err := r.ReadAll()

	if err != nil{
		exit("Error during reading file")
	}

	fmt.Println(lines)
	//current state [[5+5 10] [7+3 10] [1+1 2] [8+3 11] [1+2 3] [8+6 14] [3+1 4] [1+4 5] [5+1 6] [2+3 5] [3+3 6] [2+4 6] [5+2 7]]

	arrayOfObject := parseLines(lines)
	// current state[{5+5 10} {7+3 10} {1+1 2} {8+3 11} {1+2 3} {8+6 14} {3+1 4} {1+4 5} {5+1 6} {2+3 5} {3+3 6} {2+4 6} {5+2 7}]

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correctAnswer := 0

problemloop:
	for _,line:= range arrayOfObject {
		fmt.Println("Question is", line.question)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer
		}()


	select {
		case <-timer.C:
			fmt.Println()
			break problemloop

			//the value from channel insert to answer variable
		case answer := <-answerCh:
			if answer == line.answer {
				correctAnswer++
			}
		}
	}
	fmt.Println("You scored %d out of %d.\n",correctAnswer, len(arrayOfObject))
}

//lines [][]string is like line [[]]
func parseLines(lines [][]string)[]problem{
	arrayRes := make([]problem, len(lines))

	for i,l := range lines {
		arrayRes[i] = problem{
			question: l[0], //first column is question
			answer:   l[1], //second column is answer
		}
	}
	return arrayRes
}

func exit(message string){
	fmt.Println(message)
	os.Exit(1)
}