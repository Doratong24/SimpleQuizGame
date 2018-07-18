package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./load"
)

func main() {
	problems := load.ReadCSV()
	start(problems, 5)
}

func start(problems [][]string, second int) {
	cmdReader := bufio.NewReader(os.Stdin)
	score := 0

	for i := 0; i < len(problems); i++ {
		ans := problems[i][1]
		ques := problems[i][0]
		fmt.Println(ques)

		userInput, _ := cmdReader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if strings.Compare(ans, userInput) == 0 {
			fmt.Println("Correct")
			score++
		} else {
			fmt.Println("Incorrect")
		}

	}

	fmt.Printf("Total Score: %d\n", score)
}
