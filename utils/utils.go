package utils

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Contains(array []int, item int) bool {
	i := sort.SearchInts(array, item)
	return i < len(array) && array[i] == item
}

func ReadInts(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	return input
}
