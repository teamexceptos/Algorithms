package main

import ("fmt"
		"math")

func main()  {

	var product int = 0

	for i := 2; i < 500; i++ {
		for j := 2; j < 500; j++ {

			var m = i*i
			var n = j*j
			var sum = float64(m + n)
			var sumsq = math.Sqrt(sum)

			if(int64(sum) % int64(sumsq) == 0) {

				if(i + j + int(sumsq) == 1000){

					product = i * j * int(sumsq)
				
					break;
				}
			}
		}
	}

	fmt.Printf("%v", product)
}