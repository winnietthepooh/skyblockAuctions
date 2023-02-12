package main

import (
	"fmt"
	"skyblockAuctions/auctions"
	"time"
)

func main() {
	start := time.Now()
	auctions.DoAll()
	fmt.Println(time.Since(start))
}
