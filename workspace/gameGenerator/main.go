package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func initializePool() []int {
	var numberPool []int

	for i := 0; i < 60; i++ {
		numberPool = append(numberPool, i+1)
	}

	return numberPool
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func drawNumbers(pool []int) (numbers []int, newPool []int) {
	var game []int

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < 6; i++ {
		idx := 0
		if len(pool) > 1 {
			idx = r.Intn(len(pool) - 1)
		}

		game = append(game, pool[idx])
		pool = remove(pool, idx)
	}

	sort.Ints(game)
	return game, pool
}

func main() {
	pool := initializePool()

	var games [][]int
	for i := 0; i < 10; i++ {
		var currGame []int

		currGame, pool = drawNumbers(pool)
		games = append(games, currGame)
	}

	fmt.Println(games)
}
