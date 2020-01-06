package main

import ("fmt")

func contains(slice []int, item int) bool {
    set := make(map[int]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}

func Multiplestill20(int_v int) bool {

	var s []int

	for i := 1; i < 20; i++ {
		
		if(int_v % i == 0) {
			s = append(s, 1)

		} else {
			s = append(s, 0)
		}
	}

	if(!contains(s, 0)) {

		return true
	}
	
	return false
}


func main()  {

	var s []int

	for i := 180000000; i < 270000000; i++ {

		if(Multiplestill20(i)) {

			s = append(s, i)
		}
	}

	fmt.Printf("%v", s)
}