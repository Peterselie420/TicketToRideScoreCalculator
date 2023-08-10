package main

import (
	"golang.org/x/exp/slices"
)

// Given a network, calculate the maximum possible score based on route length
// and cards
func calculateScore(network []route) (int, []card) {
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



