package main

import (
	"fmt"
	"quiz_slice_map/check"
)

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	average, filtered := check.Checking(scores)
	fmt.Println("rata-rata : ", average, "Diatas 90 : ", filtered)
}
