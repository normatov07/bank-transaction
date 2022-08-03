package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxy"

func init() {
	rand.Seed(time.Now().UnixNano() + 1)

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

//Generate random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

//Generate a random currency code
func RandomCurrancy() string {
	var currancy []string = []string{"USD", "RUB", "SUM"}
	n := len(currancy)
	return currancy[rand.Intn(n)]
}
