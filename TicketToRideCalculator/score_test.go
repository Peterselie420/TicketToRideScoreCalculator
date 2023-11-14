package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var networkSize = 5
var network []route

func initRandomNetwork() {
	for len(network) < networkSize {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		n := r1.Intn(len(routes))
		network = append(network, routes[n])
	}
}

func BenchmarkScoreCalculationRandomNetwork(b *testing.B) {
	initData()

	b.Run("ScoreRandomNetwork", func(b *testing.B) {
		b.ResetTimer()
		initRandomNetwork()
		for i := 0; i < b.N; i++ {
			scoreLength, scoreCards, _ := calculateScore(network)
			network = nil
			fmt.Println(scoreLength + scoreCards)
		}
	})
}

func BenchmarkScoreCalculationCompleteNetwork(b *testing.B) {
	initData()

	b.Run("ScoreCompleteNetwork", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			scoreLength, scoreCards, _ := calculateScore(routes)
			fmt.Println(scoreLength + scoreCards)
		}
	})
}

func TestCheckCardInNetworkFromAndToMissing(t *testing.T) {
	card := athinaAncora

	result := checkCardInNetwork([][]city{}, card)

	if result {
		t.Errorf("Missing card was still detected")
	}
}

func TestCheckCardInNetworkFromMissing(t *testing.T) {
	card := athinaAncora
	_cityList := [][]city{{ancora}}

	result := checkCardInNetwork(_cityList, card)

	if result {
		t.Errorf("Card with missing From was still detected")
	}
}

func TestCheckCardInNetworkToMissing(t *testing.T) {
	card := athinaAncora
	_cityList := [][]city{{athina}}

	result := checkCardInNetwork(_cityList, card)

	if result {
		t.Errorf("Card with missing From was still detected")
	}
}

func TestCheckCardInNetworkFromAndToDisconnected(t *testing.T) {
	card := athinaAncora
	_cityList := [][]city{{athina}, {ancora}}

	result := checkCardInNetwork(_cityList, card)

	if result {
		t.Errorf("Card with disconnected From and To was still detected")
	}
}

func TestCheckCardInNetworkFromAndToConnected(t *testing.T) {
	card := athinaAncora
	_cityList := [][]city{{athina, ancora}}

	result := checkCardInNetwork(_cityList, card)

	if !result {
		t.Errorf("Present card was not detected")
	}
}
