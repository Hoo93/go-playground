package main

import (
    "fmt"
    "time"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
	return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
	return v
    } else {
	fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func printType(v interface{}) {
    fmt.Printf("variable : %v , Type : %T\n",v ,v)
}

func main() {
    fmt.Println("When's a Saturday?")
    today := time.Now().Weekday()
    printType(today)
    switch time.Saturday {
    case today + 0:
	fmt.Println("Today. ")
    case today + 1:
	fmt.Println("Tomorrow.")
    case today + 2:
	fmt.Println("In two days.")
    default:
	fmt.Println("Too far away.")
    } 

}


