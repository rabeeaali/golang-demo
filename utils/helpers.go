package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func initGer() {
	rand.Seed(time.Now().UnixNano())
}

//RandNumbers : make a random numbers
func RandNumbers() string {
	initGer()
	min := 1
	max := 999999
	return strconv.Itoa(rand.Intn(max-min+1) + 1)
}