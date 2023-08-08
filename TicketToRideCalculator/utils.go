package main

import (
	"fmt"
	"log"
)

// Calculate score based on length
func lengthToScore(length int) int {
	switch length {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 7
	case 6:
		return 15
	case 8:
		return 21
	}
	log.Fatal("Unknown length")
	return 0
}


// Check if two routes are connected
func routesAreConnected(route route, _route route) bool {
	return ((route.from.id == _route.from.id) || (route.from.id == _route.to.id) ||
		(route.to.id == _route.from.id) || (route.to.id == _route.to.id))
}


func printRouteSlice(routes []route) {
	for _, route := range routes {
		fmt.Print(route.from.name)
		fmt.Print(" - ")
		fmt.Print(route.to.name)
		fmt.Print(" && ")
	}
	fmt.Println("")
}

func printCards(cards []card) {
	for _, card := range cards {
		fmt.Print(card.from.name)
		fmt.Print(" - ")
		fmt.Print(card.to.name)
		fmt.Print(" _ ")
		fmt.Print(card.points)
		fmt.Print(" && ")
	}
	fmt.Println("")
}

// Given two slices, see if their content is identical (disregarding position)
// (I stole this code from somewhere no clue what it does)
func sameStringSlice(x, y []int) bool {
    if len(x) != len(y) {
        return false
    }
    // create a map of string -> int
    diff := make(map[int]int, len(x))
    for _, _x := range x {
        // 0 value for int is 0, so just increment a counter for the string
        diff[_x]++
    }
    for _, _y := range y {
        // If the string _y is not in diff bail out early
        if _, ok := diff[_y]; !ok {
            return false
        }
        diff[_y] -= 1
        if diff[_y] == 0 {
            delete(diff, _y)
        }
    }
    return len(diff) == 0
}
