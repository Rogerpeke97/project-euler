package problem6

import "math"

/*
<p>The sum of the squares of the first ten natural numbers is,</p>
$$1^2 + 2^2 + ... + 10^2 = 385.$$
<p>The square of the sum of the first ten natural numbers is,</p>
$$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
<p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
<p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
*/

func SumSquareDiff () (int){
  var sum_of_squares float64
  var square_of_sum float64
  for i := 1; i < 101; i++ {
    sum_of_squares += math.Pow(float64(i), 2)
    square_of_sum += float64(i)
  }
  return int(math.Pow(square_of_sum, 2) - sum_of_squares)
}
