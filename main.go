package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"./load"
)

func shuffle(questions [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())

	dest := make([][]string, len(questions))
	perm := rand.Perm(len(questions))
	for i, v := range perm {
		dest[i] = questions[v]
	}
	return dest
}

func main() {
	waitTime := flag.Int("wait", 30, "Question time limit")
	shuffBool := flag.Bool("shuffle", false, "Shuffle question order?")
	flag.Parse()

	problems := load.ReadCSV()
	if *shuffBool {
		problems = shuffle(problems)
	}
	start(problems, *waitTime)
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
		fmt.Println("Press ENTER to start!")

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

			case <-time.After(time.Duration(second) * time.Second):
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
