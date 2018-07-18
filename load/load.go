package load

import (
	"bufio"
	"encoding/csv"
	"os"
)

func ReadCSV() [][]string {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	all, _ := reader.ReadAll()
	return all

}
