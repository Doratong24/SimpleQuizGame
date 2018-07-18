package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"./load"
)

func main() {
	problems := load.ReadCSV()
	start(problems, 5)
}

func getInput(input chan string) {
	for {
		cmdReader := bufio.NewReader(os.Stdin)
		result, err := cmdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		result = strings.Replace(result, "\n", "", -1)
		input <- result
	}
}

func start(problems [][]string, second int) {
	input := make(chan string, 1)
	go getInput(input)
	score := 0

	for i := 0; i < len(problems); i++ {
		ans := problems[i][1]
		ques := problems[i][0]
		fmt.Println(ques)

		for {
			done := false
			fmt.Printf("Answer:")
			select {
			case i := <-input:
				if strings.Compare(ans, i) == 0 {
					fmt.Println("Correct")
					score++
					done = true
				} else {
					fmt.Println("Incorrect")
					done = true
				}

			case <-time.After(2 * time.Second):
				fmt.Println("Time out")
				done = true
			}
			if done {
				break
			}
		}
	}

	fmt.Printf("Total Score: %d/%d\n", score, len(problems))
}
