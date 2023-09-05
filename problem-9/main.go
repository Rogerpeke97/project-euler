package problem9

import (
	"math"
)

/*
<p>A Pythagorean triplet is a set of three natural numbers, $a \lt b \lt c$, for which,
$$a^2 + b^2 = c^2.$$</p>
<p>For example, $3^2 + 4^2 = 9 + 16 = 25 = 5^2$.</p>
<p>There exists exactly one Pythagorean triplet for which $a + b + c = 1000$.<br>Find the product $abc$.</p>
*/


func FindPythagoreanTriplet() ([3]int) {
  sum_should_be := 1000
  var sum_nums [3]int
  for i := 1; i < 400; i++ {
    for j := 1; j < 600; j++ {
      c := math.Pow(float64(i), 2) + math.Pow(float64(j), 2)
      c_sqrt := math.Sqrt(c)
      sum := i + j + int(c_sqrt)
      if sum == sum_should_be && i < j && c_sqrt == float64(int(c_sqrt)) {
        sum_nums[0] = i
        sum_nums[1] = j
        sum_nums[2] = int(c_sqrt)
        return sum_nums
      }
    }
  }
  return sum_nums
}
