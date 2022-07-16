package main

import "fmt"

func main() {

	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}

	res := corpFlightBookings(bookings, 6)

	fmt.Println(res)

}

func corpFlightBookings(bookings [][]int, n int) []int {
	if n <= 0 {
		return make([]int, 0)
	}

	diff := make([]int, n+1)

	for i := 0; i < len(bookings); i++ {
		increment(diff, bookings[i][0], bookings[i][1], bookings[i][2])
	}

	res := make([]int, n+1)
	res[0] = diff[0]
	for j := 1; j < len(diff); j++ {
		res[j] = res[j-1] + diff[j]
	}

	return res[1 : n+1]
}

func increment(diff []int, i, j, val int) {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
}
