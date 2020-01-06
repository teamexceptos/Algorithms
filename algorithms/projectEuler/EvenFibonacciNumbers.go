package main

import ("fmt")

func unique(intSlice []int) []int {
    keys := make(map[int]bool)
	list := []int{} 
	
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func main()  {

	var first int = 1
	var second int = 2
	
	var sum = first + second
	var totalsums int = 0

	var s []int

	for sum > second {

		s = append(s, first)
		s = append(s, second)

		first = second
		second = sum

		sum = first + second

		if ( sum > 6000000) {
			break;
		}
	}

	uniqueSlice := unique(s)

	for i := 0; i < len(uniqueSlice); i++ {

		if(uniqueSlice[i] % 2 == 0) {

			totalsums += uniqueSlice[i]
		}
	}

	fmt.Printf("%v", uniqueSlice)
	fmt.Println("totalarray: ", totalsums)
}