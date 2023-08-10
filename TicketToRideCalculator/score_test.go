package main

import (
	"math/rand"
  	"time"
    "testing"
)

var networkSize = 90
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
            calculateScore(network)
			network = nil
        }
    })
}

func BenchmarkScoreCalculationCompleteNetwork(b *testing.B) {
	initData()

    b.Run("ScoreCompleteNetwork", func(b *testing.B) {
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            calculateScore(routes)
			network = nil
        }
    })
}