package main

import ("fmt"
		"math/big")

func main()  {

	var s int = 0

	for i := 0; i < 2000000; i++ {

		var n = int64(i)
		var primeValue = big.NewInt(n).ProbablyPrime(0)

		if (primeValue) {

			if (i < 2000000) {
				s += i
				
			} else {
				break;
			}
		}
	}
	
	fmt.Printf("%v", s)
}