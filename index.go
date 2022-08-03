package main

import (
	"bank/db/util"
	"fmt"
	"time"
)

func main() {
	fmt.Println(util.RandomCurrancy())
	fmt.Println(util.RandomOwner())
	fmt.Println(util.RandomMoney())
	fmt.Println(time.Now().UnixNano())
}
