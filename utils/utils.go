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

type PairInt64 struct {
	X int64
	Y int64
}

func (p PairInt64) Sum() int64 {
	return p.X + p.Y
}

func Combinations(array []int64) []PairInt64 {
	var result []PairInt64
	for i, x := range array {
		for j, y := range array {
			if i != j {
				result = append(result, PairInt64{X: x, Y: y})
			}
		}
	}
	return result
}

func Sum(array []int64) int64 {
	var result int64 = 0
	for _, value := range array {
		result += value
	}
	return result
}

func MinAndMax(array []int64) PairInt64 {
	min := array[0]
	max := array[0]
	for _, item := range array {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}
	return PairInt64{min, max}
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

func ReadInt64s(filepath string) []int64 {
	file, err := os.Open(filepath)
	HandleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
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

type Point struct{ X, Y int }

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

var Origin = Point{0, 0}

type Grid struct {
	Geology []string
}

func (g Grid) IsTree(p Point) bool {
	i := p.X % len(g.Geology[0])
	return g.Geology[p.Y][i:i+1] == "#"
}

func (g Grid) Contains(p Point) bool {
	return p.Y >= 0 && p.Y < len(g.Geology)
}

type Passport struct {
	PassportData map[string]string
}

var requiredPassportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func (p Passport) HasRequiredFields() bool {
	for _, field := range requiredPassportFields {
		if p.PassportData[field] == "" {
			return false
		}
	}
	return true
}

func (p Passport) ValidBirthYear() bool {
	return validYear(p.PassportData["byr"], 1920, 2002)
}

func (p Passport) validIssueYear() bool {
	return validYear(p.PassportData["iyr"], 2010, 2020)
}

func (p Passport) validExpirationYear() bool {
	return validYear(p.PassportData["eyr"], 2020, 2030)
}

func validYear(year string, min int, max int) bool {
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return y >= min && y <= max
}

func (p Passport) ValidHeight() bool {
	var heightRegex = regexp.MustCompile(`(?P<value>\d+)(?P<units>[a-z]+)`)
	result := FindNamedMatches(heightRegex, p.PassportData["hgt"])
	value, err := strconv.Atoi(result["value"])
	if err != nil {
		return false
	}
	return (result["units"] == "cm" && value >= 150 && value <= 193) || (result["units"] == "in" && value >= 59 && value <= 76)
}

func (p Passport) ValidHairColour() bool {
	match, _ := regexp.MatchString("#[a-f0-9]{6}", p.PassportData["hcl"])
	return match
}

var validEyeColours = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func (p Passport) ValidEyeColour() bool {
	return validEyeColours[p.PassportData["ecl"]]
}

func (p Passport) ValidPassportId() bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", p.PassportData["pid"])
	return match
}

func (p Passport) HasValidData() bool {
	return p.HasRequiredFields() &&
		p.ValidBirthYear() &&
		p.validIssueYear() &&
		p.validExpirationYear() &&
		p.ValidHeight() &&
		p.ValidHairColour() &&
		p.ValidEyeColour() &&
		p.ValidPassportId()
}
