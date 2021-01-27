package main

import (
    "fmt"
    //"math"
)

func convert(x float64) float64{
    usd := 195.12
    return usd*x
    //return math.Round((usd*x)*100)/100
    // return math.Ceil(usd*x)
}

func main() {
    fmt.Println("Enter your usd amount needed to be converted to lkr")
    var first float64
    fmt.Scanln(&first) 
    fmt.Print("Your lkr amount is: ",convert(first)) 
}