package main

import ("fmt"
		"math/big")

func main()  {

	var s []int64
	var tractfor int64 = 1

	for len(s) < 10002 {

		trackfor += 1

		if big.NewInt(trackfor).ProbablyPrime(0) {
			s = append(s, trackfor)
		}
	}
	
	fmt.Printf("%v", len(s))
}