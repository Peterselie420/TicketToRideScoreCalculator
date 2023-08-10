package main

import (
	"golang.org/x/exp/slices"
)

// Given a network, calculate the maximum possible score based on route length
// and cards
func calculateScore(network []route) (int, int, []card) {
	var scoreLength = 0
	var scoreCards = 0
	// Calculate score based on route length
	for _, route := range network {
		scoreLength += lengthToScore(route.length)
	}
	var _cities = buildCitiesList(network)
	var _cards []card
	// Calculate score based on cards
	for _, card := range cards {
		if checkCardInNetwork(_cities, card) {
			_cards = append(_cards, card)
			scoreCards += card.points
		}
	}
	// Placeholder card
	var langsteUnit card = berlinRoma
	// Calculate score based on long cards
	for _, card := range langeUnits {
		if checkCardInNetwork(_cities, card) {
			if card.points > langsteUnit.points {
				langsteUnit = card
			}
		}
	}
	// If we found a long route that is in the network, add the score
	if langsteUnit != berlinRoma {
		_cards = append(_cards, langsteUnit)
		scoreCards += langsteUnit.points
	}
	return scoreLength, scoreCards, _cards
}

func buildCitiesList(network []route) [][]city {
	var _cities [][]city
	var routesVisited []route
	// For all routes in the network...
	for _, _route := range network {
		// If we have not already visited this route in a subset...
		if !slices.Contains(routesVisited, _route) {
			var networkCities []city
			routesVisited = append(routesVisited, _route)
			networkCities = append(networkCities, _route.from, _route.to)
			_cities = append(_cities, checkConnections(network, networkCities, routesVisited, _route))
			}
	}
	return _cities
}

// Given a network and a starting route, returns a list of connected cities
func checkConnections(network []route, networkCities []city, routesVisited []route, route route) []city {
	// For all connections of the given route... 
	for _, _route := range routeMap[route] {
		// If the network contains the connection, and we have not already seen this route...
		if slices.Contains(network, _route) && !slices.Contains(routesVisited, _route) {
			routesVisited = append(routesVisited, _route)
			if !slices.Contains(networkCities, _route.from) {
				networkCities = append(networkCities, _route.from)
			}
			if !slices.Contains(networkCities, _route.to) {
				networkCities = append(networkCities, _route.to)
			}
			checkConnections(network, networkCities, routesVisited, _route)
		}
	}
	return networkCities
}

func checkCardInNetwork(_cities [][]city, card card) bool {
	for _, _cityList := range _cities {
		if slices.Contains(_cityList, card.from) && slices.Contains(_cityList, card.to) {
			return true
		}
	}
	return false
}



