package main

import (
	"fmt"
	problem1 "project-euler/problem-1"
	problem2 "project-euler/problem-2"
	problem3 "project-euler/problem-3"
	problem4 "project-euler/problem-4"
	problem5 "project-euler/problem-5"
	problem6 "project-euler/problem-6"
	problem7 "project-euler/problem-7"
	problem8 "project-euler/problem-8"
)





func main () {
  problem1_result := problem1.ListAllNaturalNumsBelowLimit(1000)
  fmt.Printf("All natural nums below limit 1000 are %+v\n", problem1_result)

  problem2_result := problem2.EvenFibonacci()
  fmt.Printf("Even value sum is %+v\n", problem2_result)

  problem3_result := problem3.FindLargestPrimeFactor()
  fmt.Printf("Largest prime factor is %+v\n", problem3_result)
  
  highest_palindrome := problem4.LargestPalindromeProductOfThreeDigits()
  fmt.Printf("Highest palindrome of three digits is: %+v\n", highest_palindrome)

  smallest_positive_number_even_no_remainder := problem5.SmallestPositiveNumberEvenNoRemainder()
  fmt.Printf("smallest blah: %+v\n", smallest_positive_number_even_no_remainder)

  sum_square_diff := problem6.SumSquareDiff()
  fmt.Printf("sum square diff is: %+v\n", sum_square_diff)

  prime10001st := problem7.Prime10001st()
  fmt.Printf("prime10001st is: %+v\n", prime10001st)

  thirteen_adjacent_greatest_product := problem8.GreatestThirteenAdjacent()
  fmt.Printf("thirteen_adjacent_greatest_product is: %+v\n", thirteen_adjacent_greatest_product) 
}

