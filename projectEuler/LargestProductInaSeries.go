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

func main() {

	var algoDir = "/home/cosmic/Documents/allcodings/"

    file, err := os.Open(algoDir + "/Algorithms/projectEuler/LargestProductInaSeries")

    if err != nil {
        log.Fatal(err)
	}
	
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        lineStr := scanner.Text()
        // num, _ := strconv.Atoi(lineStr)
        
        fmt.Println(lineStr)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}