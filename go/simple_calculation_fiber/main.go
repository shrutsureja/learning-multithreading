package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		p := fmt.Println
		start := time.Now()
		MIL := 100000000
		var totalPrimeNumbers int
		for i := 1; i <= MIL; i++ {
			if isPrime(i) {
				totalPrimeNumbers++
			}
		}
		fmt.Println(totalPrimeNumbers)
		end := time.Now()
		p(end.Sub(start))
		return nil
	})

	app.Listen(":3000")
}