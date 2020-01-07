package main

import ("fmt")

func palindromechecker(int_v int) bool {

	var v, remainder, temp int
    var reverse int = 0
 
	temp = int_v
	v = int_v
 
    for {
        remainder = v%10
        reverse = reverse*10 + remainder
        v /= 10
 
        if(v == 0 ){
            break;
        }
	}
	
	if(temp == reverse) {
		return true
	}
	return false
}

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

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main()  {

	var s []int

	for i := 100; i < 1000; i++ {
			
		for j := 100; j < 1000; j++ {
		
			var mul int = i * j

			if (palindromechecker(mul)){
				s = append(s, mul)
			}
		}
	}

	uniqueSlice := unique(s)
	min, max := findMinAndMax(uniqueSlice)

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

	
}