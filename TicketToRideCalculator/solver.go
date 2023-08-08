package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/slices"
)

// Amount of trains available (in the real game this is 46?)
// for 7 trains, the solver runs in a reasonable time (few minutes :D)
// where it should find a maximum score of: 18 (brest to marseille with 1 card)
const trainsMax int = 7

var score int = 0
var consideredNetworks [][]int
var langeUnits []card

// Initialize variables and call solve method
func setupSolver(routes []route, cards []card, longCards []card, routeMap map[route][]route) {
	var trainsUsed int = 0
	var network []route
	langeUnits = longCards
	barLength := len(routes)
	bar := progressbar.Default(int64(barLength))
	solve(routes, cards, routeMap, network, trainsUsed, bar)
}

// Given a network, calculate the maximum possible score based on route length
// and cards
func calculateScore(network []route, routeMap map[route][]route) (int, []card) {
	var localScore = 0
	// Calculate score based on route length
	for _, route := range network {
		localScore += lengthToScore(route.length)
	}
	var _cards []card
	// Calculate score based on cards
	for _, card := range cards {
		if cardInNetwork(network, routeMap, card) {
			_cards = append(_cards, card)
			localScore += card.points
		}
	}
	// Placeholder card
	var langsteUnit card = berlinRoma
	// Calculate score based on long cards
	for _, card := range langeUnits {
		if cardInNetwork(network, routeMap, card) {
			if card.points > langsteUnit.points {
				langsteUnit = card
			}
		}
	}
	// If we found a long route that is in the network, add the score
	if langsteUnit != berlinRoma {
		_cards = append(_cards, langsteUnit)
		localScore += langsteUnit.points
	}
	return localScore, _cards
}

// Function to check if a given card is valid for a given network
// a.k.a. the cities on the card are connected in the network
func cardInNetwork(network []route, routeMap map[route][]route, card card) bool {
	var routesVisited []route
	// For all routes in the network...
	for _, _route := range network {
		// If we have not already visited this route in a subset...
		if !slices.Contains(routesVisited, _route) {
			var currentNetworkSubset []route
			var to bool
			var from bool
			routesVisited = append(currentNetworkSubset, _route)
			// If the network contains a connection between the cities on the card
			if checkConnections(network, currentNetworkSubset, routesVisited, routeMap, _route, from, to, card) {
				return true
			}
		}
	}
	return false
}

// Given a network, recursively check if there exists a connection between the cities
// on the given card, we do this by starting with a single route, and extending the route 
// network by adding a single route at a time, keeping track of which routes / subsets we consider
// as to not calculate unnessecary. If at any time, a route in the considered network contains either
// one of the cities on the card, we flag a boolean, if both booleans are flagged, the network contains
// the card
func checkConnections(network []route, currentNetworkSubset []route, routesVisited []route, routeMap map[route][]route, route route, from bool, to bool, card card) bool {
	// For all connections of the given route... 
	for _, _route := range routeMap[route] {
		// If the network contains the connection, and we have not already seen this route in a subset...
		if slices.Contains(network, _route) && !slices.Contains(currentNetworkSubset, _route) {
			routesVisited = append(routesVisited, _route)
			currentNetworkSubset = append(currentNetworkSubset, _route)
			// Check if the card.from city is in the network (the currently considered route)
			if _route.from.id == card.from.id || _route.to.id == card.from.id {
				from = true
			}
			// Check if the card.to city is in the network (the currently considered route)
			if _route.from.id == card.to.id || _route.to.id == card.to.id {
				to = true
			}
			// If both card.from and card.to are in the network, our work is done
			if from && to {
				return true
			}
			return checkConnections(network, currentNetworkSubset, routesVisited, routeMap, _route, from, to, card)
		}
	}
	return false
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
func solve(routes []route, cards []card, routeMap map[route][]route, network []route, trainsUsed int, bar *progressbar.ProgressBar) {
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
				var currentScore, _cards = calculateScore(network, routeMap)
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
			solve(routes, cards, routeMap, network, trainsUsed, bar)
			// The search with the considered route did not yield succesful, so
			// remove the latest entry (recursive magic)
			network = network[:len(network)-1]
			trainsUsed -= route.length
		}
	}
}


