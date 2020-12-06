package day05

import (
	"fmt"
	"sort"
)

type Seat struct {
	row    int
	column int
}

func (s Seat) Id() int {
	return s.row*8 + s.column
}

func parseSeat(seatCode string) Seat {
	return Seat{
		row:    calculateRow(seatCode),
		column: calculateColumn(seatCode),
	}
}

func calculateRow(seatCode string) int {
	low := 0
	high := 127
	for _, s := range seatCode[:7] {
		switch s {
		case 'F':
			high = (low + high - 1) / 2
		case 'B':
			low = (low + high + 1) / 2
		}
	}
	return low
}

func calculateColumn(seatCode string) int {
	low := 0
	high := 7
	for _, s := range seatCode[len(seatCode)-3:] {
		switch s {
		case 'L':
			high = (low + high - 1) / 2
		case 'R':
			low = (low + high + 1) / 2
		}
	}
	return low
}

func solvePart1(input []string) int {
	maxSeatId := 0
	for _, code := range input {
		seat := parseSeat(code)
		seatId := seat.Id()
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return maxSeatId
}

func solvePart2(input []string) int {
	seats := []int{}
	for _, code := range input {
		seat := parseSeat(code)
		seatId := seat.Id()
		seats = append(seats, seatId)
	}
	sort.Ints(seats)

	prev := 0
	maybe := 0
	for _, seatId := range seats {
		if seatId-prev == 2 {
			fmt.Printf("Could be %v\n", seatId-1)
			maybe = seatId - 1
		}
		prev = seatId
	}
	return maybe
}
