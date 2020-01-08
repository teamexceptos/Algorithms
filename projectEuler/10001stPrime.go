package main

import ("fmt"
		"math/big")

func main()  {

	var s []int64
	var tractfor int64 = 1

	for len(s) < 10002 {

		tractfor += 1

		if big.NewInt(tractfor).ProbablyPrime(0) {
			s = append(s, tractfor)
		}
	}
	
	fmt.Printf("%v", len(s))
}