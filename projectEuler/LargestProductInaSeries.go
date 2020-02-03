package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func AdjacentProduct(index, total int64, digits []int64) int64 {

    product := int64(1)
    for i := int64(0); i < total; i++ {
            product *= digits[index+i]
    }

    return product
}

func main() {

    var algoDir = "/home/cosmic/Documents/allcodings/"
    var digits []int64

    file, err := os.Open(algoDir + "/Algorithms/projectEuler/LargestProductInaSeries")

    if err != nil {
        log.Fatal(err)
	}
	
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        lineStr := scanner.Text()
        for _, runeValue := range lineStr {

            i, err := strconv.ParseInt(string(runeValue), 10, 0)
            if err != nil {
                    panic(err)
            }
            digits = append(digits, i)
        } 
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var max, index int64 = 0, 0
    for i := int64(0); i < int64(len(digits)-13); i++ {
        product := AdjacentProduct(i, 13, digits)
        if product > max {
            max = product
            index = i
        }
    }

    fmt.Print("Index at ")
    fmt.Println(index)

    for i := int64(0); i < 13; i++ {
        fmt.Print(digits[index+i])
        if i != 12 {
            fmt.Print(" * ")
        }
    }
    fmt.Print(" = ")
    fmt.Println(max)
}