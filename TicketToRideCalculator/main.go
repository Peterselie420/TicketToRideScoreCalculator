package main

//go:generate goversioninfo

func main() {
	routes, cards, longCards, routeMap = initData()
	setupSolver(routes, cards, longCards, routeMap)
}
