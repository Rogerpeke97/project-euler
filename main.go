package main

import (
	"fmt"
	problem1 "project-euler/problem-1"
	problem2 "project-euler/problem-2"
	problem3 "project-euler/problem-3"
)





func main () {
  problem1_result := problem1.ListAllNaturalNumsBelowLimit(1000)
  fmt.Printf("All natural nums below limit 1000 are %+v\n", problem1_result)

  problem2_result := problem2.EvenFibonacci()
  fmt.Printf("Even value sum is %+v\n", problem2_result)

  problem3_result := problem3.FindLargestPrimeFactor()
  fmt.Printf("Largest prime factor is %+v\n", problem3_result)
}

