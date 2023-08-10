package main

import (
	"log"
	"fmt"
    "github.com/schollz/progressbar/v3"
)

/**
* Contains and builds data to use in solver
*/

// City Constants
const (
    Lisboa int = iota
    Cadiz
    Madrid
    Barcelona
    Pamplona
    Brest
    Dieppe
    Paris
    London
    Edinburgh
    Amsterdam
    Bruxelles
    Essen
    Frankfurt
    Munchen
    Zurich
    Marseille
    Roma
    Venezia
    Wien
    Ancora
    Berlin
    Kobenhavn
    Stockholm
    Brindisi
    Zagrab
    Budapest
    Palermo
    Athina
    Sofia
    Sarajevo
    Smyrna
    Constantinople
    Erzrum
    Bucuristi
    Sevastopol
    Sochi
    Rostov
    Kharkov
    Kyiv
    Moskva
    Smolensk
    Wilno
    Danzic
    Warszawa
    Riga
    Petrograd
    
)

// Color Constants
const (
	Purple int = iota
	Blue
	Orange
	Yellow
	Black
	White
	Red
	Green
	Grey
    Grey2
)

type city struct {
	id   int
	name string
}

type route struct {
	from     city
	to       city
	length   int
	color    int
	locos    int
	isTunnel bool
    id       int
}

type card struct {
    from city
    to city
    points int
    long bool
}

// Cities
var cities []city
var lisboa = city{Lisboa, "Lisbon"}
var cadiz = city{Cadiz, "Cadiz"}
var madrid = city{Madrid, "Madrid"}
var barcelona = city{Barcelona, "Barcelona"}
var pamplona = city{Pamplona, "Pamplona"}
var brest = city{Brest, "Brest"}
var dieppe = city{Dieppe, "Dieppe"}
var paris = city{Paris, "Paris"}
var london = city{London, "London"}
var edinburgh = city{Edinburgh, "Edinburgh"}
var amsterdam = city{Amsterdam, "Amsterdam"}
var bruxelles = city{Bruxelles, "Bruxelles"}
var essen = city{Essen, "Essen"}
var frankfurt = city{Frankfurt, "Frankfurt"}
var munchen = city{Munchen, "Munchen"}
var zurich = city{Zurich, "Zurich"}
var marseille = city{Marseille, "Marseille"}
var roma = city{Roma, "Roma"}
var venezia = city{Venezia, "Venezia"}
var wien = city{Wien, "Wien"}
var ancora = city{Ancora, "Ancora"}
var berlin = city{Berlin, "Berlin"}
var kobenhavn = city{Kobenhavn, "Kobenhavn"}
var stockholm = city{Stockholm, "Stockholm"}
var brindisi = city{Brindisi, "Brindisi"}
var zagrab = city{Zagrab, "Zagrab"}
var budapest = city{Budapest, "Budapest"}
var palermo = city{Palermo, "Palermo"}
var athina = city{Athina, "Athina"}
var sofia = city{Sofia, "Sofia"}
var sarajevo = city{Sarajevo, "Sarajevo"}
var smyrna = city{Smyrna, "Smyrna"}
var constantinople = city{Constantinople, "Constantinople"}
var erzrum = city{Erzrum, "Erzrum"}
var bucuristi = city{Bucuristi, "Bucuristi"}
var sevastopol = city{Sevastopol, "Sevastopol"}
var sochi = city{Sochi, "Sochi"}
var rostov = city{Rostov, "Rostov"}
var kharkov = city{Kharkov, "Kharkov"}
var kyiv = city{Kyiv, "Kyiv"}
var moskva = city{Moskva, "Moskva"}
var smolensk = city{Smolensk, "Smolensk"}
var wilno = city{Wilno, "Wilno"}
var danzic = city{Danzic, "Danzic"}
var warszawa = city{Warszawa, "Warszawa"}
var riga = city{Riga, "Riga"}
var petrograd = city{Petrograd, "Petrograd"}


//Append all cities in a single list
func initCities() []city {
	cities = append(cities, lisboa, cadiz, madrid, barcelona, pamplona)
	fmt.Printf("Found %d cities\n", len(cities))
	for _, city := range cities {
		fmt.Print(city.name + " ")
	}
	fmt.Println("")
    return cities
}

//routes
var routes []route
var lisboaMadrid = route{lisboa,  madrid,  3,  Purple,  0,  false, 1}
var cadizLisboa = route{cadiz,  lisboa,  2,  Blue,  0,  false, 2}
var cadizMadrid = route{cadiz,  madrid,  3,  Orange,  0,  false, 3}
var barcelonaMadrid = route{barcelona,  madrid,  2,  Yellow,  0,  false, 4}
var madridPamplona = route{madrid,  pamplona,  3,  Black,  0,  true, 5}
var pamplonaMadrid = route{pamplona,  madrid,  3,  White,  0,  true, 6}
var barcelonaPamplona = route{barcelona,  pamplona,  2,  Grey,  0,  true, 7}
var brestPamplona = route{brest,  pamplona,  4,  Purple,  0,  false, 8}
var brestDieppe = route{brest,  dieppe,  2,  Orange,  0,  false, 9}
var dieppeLondon = route{dieppe,  london,  2,  Grey,  1,  false, 10}
var londonDieppe = route{london,  dieppe,  2,  Grey2,  1,  false, 11}
var edinburghLondon = route{edinburgh,  london,  4,  Black,  0,  false, 12}
var londonEdinburgh = route{london,  edinburgh,  4,  Orange,  0,  false, 13}
var amsterdamLondon = route{amsterdam,  london,  2,  Grey,  2,  false, 14}
var amsterdamBruxeleles = route{amsterdam,  bruxelles,  1,  Black,  0,  false, 15}
var bruxellesDieppe = route{bruxelles,  dieppe,  2,  Green,  0,  false, 16}
var dieppeParis = route{dieppe,  paris,  1,  Purple,  0,  false, 17}
var barcelonaMarseille = route{barcelona,  marseille,  4,  Grey,  0,  false, 18}
var marseillePamplona = route{marseille,  pamplona,  4,  Red,  0,  false, 19}
var marseilleParis = route{marseille,  paris,  4,  Grey,  0,  false, 20}
var parisZurich = route{paris,  zurich,  3,  Grey,  0,  true, 21}
var parisPamplona = route{paris,  pamplona,  4,  Blue,  0,  false, 22}
var pamplonaParis = route{pamplona,  paris,  4,  Green,  0,  false, 23}
var brestParis = route{brest,  paris,  3,  Black,  0,  false, 24}
var bruxellesParis = route{bruxelles,  paris,  2,  Yellow,  0,  false, 25}
var parisBruxelles = route{paris,  bruxelles,  2,  Red,  0,  false, 26}
var amsterdamEssen = route{amsterdam,  essen,  3,  Yellow,  0,  false, 27}
var amsterdamFrankfurt = route{amsterdam,  frankfurt,  2,  White,  0,  false, 28}
var bruxellesFrankfurt = route{bruxelles,  frankfurt,  2,  Blue,  0,  false, 29}
var frankfurtParis = route{frankfurt,  paris,  3,  White,  0,  false, 30}
var parisFrankfurt = route{paris,  frankfurt,  3,  Orange,  0,  false, 31}
var essenFrankfurt = route{essen,  frankfurt,  2,  Green,  0,  false, 32}
var marseilleZurich = route{marseille,  zurich,  2,  Purple,  0,  true, 33}
var frankfurtMunchen = route{frankfurt,  munchen,  2,  Purple,  0,  false, 34}
var munchenZurich = route{munchen,  zurich,  2,  Yellow,  0,  true, 35}
var veneziaZurich = route{venezia,  zurich,  2,  Green,  0,  true, 36}
var munchenVenezia = route{munchen,  venezia,  2,  Blue,  0,  true, 37}
var marseilleRoma = route{marseille,  roma,  4,  Grey,  0,  true, 38}
var romaVenezia = route{roma,  venezia,  2,  Black,  0,  false, 39}
var berlinEssen = route{berlin,  essen,  2,  Blue,  0,  false, 40}
var berlinFrankfurt = route{berlin,  frankfurt,  3,  Black,  0,  false, 41}
var frankfurtBerlin = route{frankfurt,  berlin,  3,  Red,  0,  false, 42}
var essenKobenhavn = route{essen,  kobenhavn,  3,  Grey,  1,  false, 43}
var kobenhavnEssen = route{kobenhavn,  essen,  3,  Grey2,  1,  false, 44}
var kobenhavnStockholm = route{kobenhavn,  stockholm,  3,  Yellow,  0,  false, 45}
var stockholmKobenhavn = route{stockholm,  kobenhavn,  3,  White,  0,  false, 46}
var petrogradStockholm = route{petrograd,  stockholm,  8,  Grey,  0,  true, 47}
var munchenWien = route{munchen,  wien,  3,  Orange,  0,  false, 48}
var veneziaZagrab = route{venezia,  zagrab,  2,  Grey,  0,  false, 49}
var brindisiRoma = route{brindisi,  roma,  2,  White,  0,  false, 50}
var palermoRoma = route{palermo,  roma,  4,  Grey,  1,  false, 51}
var brindisiPalermo = route{brindisi,  palermo,  3,  Grey,  1,  false, 52}
var athinaBrindisi = route{athina,  brindisi,  4,  Grey,  1,  false, 53}
var palermoSmyrna = route{palermo,  smyrna,  6,  Grey,  2,  false, 54}
var athinaSmyrna = route{athina,  smyrna,  2,  Grey,  1,  false, 55}
var ancoraSmyrna = route{ancora,  smyrna,  3,  Orange,  0, true, 56}
var constantinopleSmyrna = route{constantinople,  smyrna,  2,  Grey,  0,  true, 57}
var ancoraConstantinople = route{ancora,  constantinople,  2,  Grey,  0,  true, 58}
var ancoraErzrum = route{ancora,  erzrum,  3,  Black,  0,  false, 59}
var athinaSarajevo = route{athina,  sarajevo,  4,  Green,  0,  false, 60}
var sarajevoZagrab = route{sarajevo,  zagrab,  3,  Red,  0,  false, 61}
var budapestWien = route{budapest,  wien,  1,  Red,  0,  false, 62}
var wienBudapest = route{wien,  budapest,  1,  White,  0,  false, 63}
var budapestZagrab = route{budapest,  zagrab,  2,  Orange,  0,  false, 64}
var budapestSarajevo = route{budapest,  sarajevo,  3,  Purple,  0,  false, 65}
var sarajevoSofia = route{sarajevo,  sofia,  2,  Grey,  0,  true, 66}
var bucuristiBudapest = route{bucuristi,  budapest,  4,  Grey,  0,  true, 67}
var bucuristiSofia = route{bucuristi,  sofia,  2,  Grey,  0,  false, 68}
var budapestKyiv = route{budapest,  kyiv,  6,  Grey,  0,  true, 69}
var bucuristiKyiv = route{bucuristi,  kyiv,  4,  Grey,  0,  false, 70}
var bucuristiConstantinople = route{bucuristi,  constantinople,  3,  Yellow,  0,  false, 71}
var bucuristiSevastopol = route{bucuristi,  sevastopol,  4,  White,  0,  false, 72}
var constantinopleSevastopol = route{constantinople,  sevastopol,  4,  Grey,  2,  false, 73}
var erzrumSochi = route{erzrum,  sochi,  3,  Red,  0,  false, 74}
var sevastopolSochi = route{sevastopol,  sochi,  2,  Grey,  1,  false, 75}
var rostovSochi = route{rostov,  sochi,  2,  Grey,  0,  false, 76}
var rostovSevastopol = route{rostov,  sevastopol,  4,  Grey,  0,  false, 77}
var kharkovRostov = route{kharkov,  rostov,  2,  Green,  0,  false, 78}
var constantinopleSofia = route{constantinople,  sofia,  3,  Blue,  0,  false, 79}
var kharkovKyiv = route{kharkov,  kyiv,  4,  Grey,  0,  false, 80}
var kharkovMoskva = route{kharkov,  moskva,  4,  Grey,  0,  false, 81}
var moskvaSmolensk = route{moskva,  smolensk,  2,  Orange,  0,  false, 82}
var kyivSmolensk = route{kyiv,  smolensk,  3,  Red,  0,  false, 83}
var kyivWilno = route{kyiv,  wilno,  2,  Grey,  0,  false, 84}
var smolenskWilno = route{smolensk,  wilno,  3,  Yellow,  0,  false, 85}
var moskvaPetrograd = route{moskva,  petrograd,  4,  White,  0,  false, 86}
var petrogradWilno = route{petrograd,  wilno,  4,  Blue,  0,  false, 87}
var petrogradRiga = route{petrograd,  riga,  4,  Grey,  0,  false, 88}
var rigaWilno = route{riga,  wilno,  4,  Green,  0,  false, 89}
var danzicRiga = route{danzic,  riga,  3,  Black,  0,  false, 90}
var danzicWarszawa = route{danzic,  warszawa,  2,  Grey,  0,  false, 91}
var warszawaWilno = route{warszawa,  wilno,  3,  Red,  0,  false, 92}
var kyivWarszawa = route{kyiv,  warszawa,  4,  Grey,  0,  false, 93}
var warszawaWien = route{warszawa,  wien,  4,  Blue,  0,  false, 94}
var berlinDanzic = route{berlin,  danzic,  4,  Grey,  0,  false, 95}
var berlinWarszawa = route{berlin,  warszawa,  4,  Purple,  0,  false, 96}
var warszawaBerlin = route{warszawa,  berlin,  4,  Yellow,  0,  false, 97}

// Append all routes in a single list
func initRoutes() []route {
	routes = append(routes, lisboaMadrid, cadizLisboa, cadizMadrid, barcelonaMadrid, madridPamplona, pamplonaMadrid, barcelonaPamplona, brestPamplona, brestDieppe, dieppeLondon, londonDieppe, edinburghLondon, londonEdinburgh, amsterdamLondon, amsterdamBruxeleles, bruxellesDieppe, dieppeParis, barcelonaMarseille, marseillePamplona, marseilleParis, parisZurich, parisPamplona, pamplonaParis, brestParis, bruxellesParis, parisBruxelles, amsterdamEssen, amsterdamFrankfurt, bruxellesFrankfurt, frankfurtParis, parisFrankfurt, essenFrankfurt, marseilleZurich, frankfurtMunchen, munchenZurich, veneziaZurich, munchenVenezia, marseilleRoma, romaVenezia, berlinEssen, berlinFrankfurt, frankfurtBerlin, essenKobenhavn, kobenhavnEssen, kobenhavnStockholm, stockholmKobenhavn, petrogradStockholm, munchenWien, veneziaZagrab, brindisiRoma, palermoRoma, brindisiPalermo, athinaBrindisi, palermoSmyrna, athinaSmyrna, ancoraSmyrna, constantinopleSmyrna, ancoraConstantinople, ancoraErzrum, athinaSarajevo, sarajevoZagrab, budapestWien, wienBudapest, budapestZagrab, budapestSarajevo, sarajevoSofia, bucuristiBudapest, bucuristiSofia, budapestKyiv, bucuristiKyiv, bucuristiConstantinople, bucuristiSevastopol, constantinopleSevastopol, erzrumSochi, sevastopolSochi, rostovSochi, rostovSevastopol, kharkovRostov, constantinopleSofia, kharkovKyiv, kharkovMoskva, moskvaSmolensk, kyivSmolensk, kyivWilno, smolenskWilno, moskvaPetrograd, petrogradWilno, petrogradRiga, rigaWilno, danzicRiga, danzicWarszawa, warszawaWilno, kyivWarszawa, warszawaWien, berlinDanzic, berlinWarszawa, warszawaBerlin)
    fmt.Printf("Found %d routes\n", len(routes))
	fmt.Println("")
    checkRoutes()
    constructRouteMap(routes)
    return routes
}

// Sanity check to see if there are no duplicate routes
func checkRoutes() {
    barLength := len(routes) * len(routes)
	bar := progressbar.Default(int64(barLength))
	for _, route := range routes {
		for _, _route := range routes {
			if route == _route {
                bar.Add(1)
				continue
			}
			if (((route.from.name == _route.from.name) || (route.to.name == _route.from.name)) && ((route.from.name == _route.to.name) || (route.to.name == _route.to.name))) {
				if route.color == _route.color && route.length == _route.length {
					fmt.Println("Route from " + route.from.name + " to " + route.to.name + " is identical to route from " +
						_route.from.name + " to " + _route.to.name)
					log.Fatal("Duplicate route")
				}
			}
            bar.Add(1)
		}
	}
}

// Construct a map such that for each route all connected routes are instantly obtainable
var routeMap = make(map[route][]route)
func constructRouteMap(routes []route) {
    for _, routeKey := range routes {
        var routeConnections []route
        for _, route := range routes {
            if (routeKey == route) {
                continue
            }
            if (routesAreConnected(route, routeKey)) {
                routeConnections = append(routeConnections, route)
            }
        }
        routeMap[routeKey] = routeConnections
    }
    for key, value := range routeMap {
        fmt.Print("Route ")
        fmt.Print(key)
        fmt.Print(" is connected to: //{")
        printRouteSlice(value)
        fmt.Println("")
    }
}

// Cards
var cards []card
var athinaAncora = card{athina, ancora, 5, false}
var budapestSofia = card{budapest, sofia, 5, false}
var frankfurtKobenhavn = card{frankfurt, kobenhavn, 5, false}
var rostovErzrum = card{rostov, erzrum, 5, false}
var sofiaSmyrna = card{sofia, smyrna, 5, false}
var kyivPetrograd = card{kyiv, petrograd, 6, false}
var zurichBridinsi = card{zurich, brindisi, 6, false}
var zurichBudapest = card{zurich, budapest, 6, false}
var warszawaSmolensk = card{warszawa, smolensk, 6, false}
var zagrabBridinsi = card{zagrab, brindisi, 6, false}
var parisZagreb = card{paris, zagrab, 7, false}
var brestMarseille = card{brest, marseille, 7, false}
var londonBerlin = card{london, berlin, 7, false}
var edinburghParis = card{edinburgh, paris, 7, false}
var amsterdamPamplona = card{amsterdam, pamplona, 7, false}
var romaSmyrna = card{roma, smyrna, 8, false}
var palermoConstantinople = card{palermo, constantinople, 8, false}
var sarajevoSevastopol = card{sarajevo, sevastopol, 8, false}
var madridDieppe = card{madrid, dieppe, 8, false}
var barcelonaBruxelles = card{barcelona, bruxelles, 8, false}
var parisWien = card{paris, wien, 8, false}
var barcelonaMunchen = card{barcelona, munchen, 8, false}
var brestVenezia = card{brest, venezia, 8, false}
var smolenskRostov = card{smolensk, rostov, 8, false}
var marseilleEssen = card{marseille, essen, 8, false}
var kyivSochi = card{kyiv, sochi, 8, false}
var madridZurich = card{madrid, zurich, 8, false}
var berlinBucuresti = card{berlin, bucuristi, 8, false}
var bruxellesDanzic = card{bruxelles, danzic, 9, false}
var berlinRoma = card{berlin, roma, 9, false}
var ancoraKharkov = card{ancora, kharkov, 10, false}
var rigaBucuresti = card{riga, bucuristi, 10, false}
var essenKyiv = card{essen, kyiv, 10, false}
var veneziaConstantinople = card{venezia, constantinople, 10, false}
var londonWien = card{london, wien, 10, false}
var athinaWilno = card{athina, wilno, 11, false}
var stockholmWien = card{stockholm, wien, 11, false}
var berlinMoskva = card{berlin, moskva, 12, false}
var amsterdamWilno = card{amsterdam, wilno, 12, false}
var frankfurtSmolensk = card{frankfurt, smolensk, 13, false}

// Long Cards
var longCards []card
var lisboaDanzic = card{lisboa, danzic, 20, true}
var brestPetrogard = card{brest, petrograd, 20, true}
var palermoMoskva = card{palermo, moskva, 20, true}
var kobenhavnErzurum = card{kobenhavn, erzrum, 21, true}
var edinburghAthina = card{edinburgh, athina, 21, true}
var cadizStockholm = card{cadiz, stockholm, 21, true}

func initCards() ([]card, []card) {
	cards = append(cards, athinaAncora, budapestSofia, frankfurtKobenhavn, rostovErzrum, sofiaSmyrna, kyivPetrograd, zurichBridinsi, zurichBudapest, warszawaSmolensk, zagrabBridinsi, parisZagreb, brestMarseille, londonBerlin, edinburghParis, amsterdamPamplona, romaSmyrna, palermoConstantinople, sarajevoSevastopol, madridDieppe, barcelonaBruxelles, parisWien, barcelonaMunchen, brestVenezia, smolenskRostov, marseilleEssen, kyivSochi, madridZurich, berlinBucuresti, bruxellesDanzic, berlinRoma, ancoraKharkov, rigaBucuresti, essenKyiv, veneziaConstantinople, londonWien, athinaWilno, stockholmWien, berlinMoskva, amsterdamWilno, frankfurtSmolensk)
    longCards = append(longCards, lisboaDanzic, brestPetrogard, palermoMoskva, kobenhavnErzurum, edinburghAthina, cadizStockholm)
    fmt.Printf("Found %d cards\n", len(cards))
	fmt.Println("")
    return cards, longCards
}


func initData() ([]route, []card, []card, map[route][]route) {
    cities = initCities()
	routes = initRoutes()
    cards, longCards = initCards()
    return routes, cards, longCards, routeMap
}
