package main

import ("fmt")

func main()  {

	var sumofTheSquares int = 0
	var Sum int = 0


	for i := 1; i < 101; i++ {

		sumofTheSquares += i * i
		Sum += i
	}

	var neededDiff = Sum * Sum - sumofTheSquares

	fmt.Printf("diff: %d", neededDiff)
	
}