package main

import (
	"algo/secondb"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(4096)
	fmt.Println(secondb.Secondb(arr))
}