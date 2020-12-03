package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Contains(array []int, item int) bool {
	i := sort.SearchInts(array, item)
	return i < len(array) && array[i] == item
}

func ReadInts(filepath string) []int {
	file, err := os.Open(filepath)
	HandleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		HandleError(err)
		input = append(input, i)
	}
	return input
}

func ReadStrings(filepath string) []string {
	file, err := os.Open(filepath)
	HandleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		s := scanner.Text()
		input = append(input, s)
	}
	return input
}

func FindNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Password struct {
	Min  int
	Max  int
	Char string
	Text string
}

func (p Password) IsValid() bool {
	charCount := strings.Count(p.Text, p.Char)
	return charCount >= p.Min && charCount <= p.Max
}

func (p Password) IsValidPosition() bool {
	firstPositionMatch := p.Text[p.Min-1:p.Min] == p.Char
	secondPositionMatch := p.Text[p.Max-1:p.Max] == p.Char
	return firstPositionMatch != secondPositionMatch
}

var passwordRegex = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>[a-z]): (?P<text>[a-z]+)`)

func ParsePassword(line string) Password {
	result := FindNamedMatches(passwordRegex, line)
	min, err := strconv.Atoi(result["min"])
	HandleError(err)
	max, err := strconv.Atoi(result["max"])
	HandleError(err)
	return Password{
		Min:  min,
		Max:  max,
		Char: result["char"],
		Text: result["text"],
	}
}
