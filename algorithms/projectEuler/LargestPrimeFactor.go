package main

import ("fmt"
		"math/big")

func main()  {

	var s []int

	for i := 0; i < 10000; i++ {

		var n = int64(i)

		if big.NewInt(n).ProbablyPrime(0) {

			if(600851475143 % i == 0) { 
				s = append(s, i)
			}
		}
	}
	
	fmt.Printf("%v", s)
}