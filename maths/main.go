package maths

import (
	"math/rand"
	"time"
)

func GetRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min

}
func ProgressBar(l float64) string {
	str := ""
	i := 0
	for {
		i++
		if i == 10 {
			break
		}
		if l > float64(i) {
			str += "█"
		} else {
			str += "░"
		}
	}
	return str
}
