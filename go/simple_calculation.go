package main

import (
	"fmt"
	"math"
	"net/http"
	_ "net/http/pprof" // blank import for side effects
)

func isPrime(num int) bool {
	// for i:= 2 ; i*i <= num; i++ {
	// 	if num%i==0 {
	// 		return false
	// 	}
	// }
	// return true
	if num < 2 {
		return false
	}
	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	 go func() {
        fmt.Println(http.ListenAndServe("localhost:6060", nil))
    }()
	MIL := 100000000
	var totalPrimeNumbers int

	for i := 1; i <= MIL; i++ {
		if isPrime(i) {
			totalPrimeNumbers++
		}
	}
	fmt.Println(totalPrimeNumbers)
}