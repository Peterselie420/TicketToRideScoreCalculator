package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

// Amount of trains available (in the real game this is 46?)
// for 7 trains, the solver runs in a reasonable time (few minutes :D)
// where it should find a maximum score of: 18 (brest to marseille with 1 card)
const trainsMax int = 6

var score int = 0
var consideredNetworks [][]int
var langeUnits []card

// Initialize variables and call solve method
func setupSolver() {
	var trainsUsed int = 0
	var network []route
	langeUnits = longCards
	solve(network, trainsUsed)
}

// Function to check if we have already considered a permutation of a network;
// say our network is Lisboa-Cadiz-Madrid, then Madrid-Lisboa-Cadiz does not
// need to be calculated
func proposedNetworkAlreadyConsidered(proposedNetwork []int) bool {
	if len(consideredNetworks) == 0 {
		return false
	}
	for _, network := range consideredNetworks {
		if !sameStringSlice(proposedNetwork, network) {
			return false
		}
	}
	return true
}

// Recursively consider each subset of routes, given the max number of trains to use
func solve(network []route, trainsUsed int) {
	// For each route...
	for _, route := range routes {
		var proposedNetwork []int
		// Instantiate the network we currently consider
		for _, _route := range network {
			proposedNetwork = append(proposedNetwork, _route.id)
		}
		proposedNetwork = append(proposedNetwork, route.id)
		// If the network does not already contain the route, and we do not exceed the train constraind by 'playing' this route
		// and we have not already considered a permutation of this network...
		if !slices.Contains(network, route) && trainsUsed+route.length <= trainsMax && !proposedNetworkAlreadyConsidered(proposedNetwork) {
			network = append(network, route)
			consideredNetworks = append(consideredNetworks, proposedNetwork)
			trainsUsed += route.length

			// Since the biggest route is 8 trains, calculating the score when
			// more than 8 trains are availible is a waste, since a higher scoring network
			// must exist

			// If we have 8 or less trains left...
			if trainsUsed > trainsMax-8 {
				var currentScore, _cards = calculateScore(network)
				// If the calculated score is higher than the highest we found yet...
				if currentScore > score {
					fmt.Print("Current Network: //")
					printRouteSlice(network)
					score = currentScore
					fmt.Print("Score: ")
					fmt.Println(score)
					fmt.Print("Cards: ")
					fmt.Println(len(_cards))
					printCards(_cards)
					fmt.Println("")
				}
			}
			// Recursively continue our search
			solve(network, trainsUsed)
			// The search with the considered route did not yield succesful, so
			// remove the latest entry (recursive magic)
			network = network[:len(network)-1]
			trainsUsed -= route.length
		}
	}
}


