package problem1

/*
<p>If we list all the natural numbers below $10$ that are multiples of $3$ or $5$, we get $3, 5, 6$ and $9$. The sum of these multiples is $23$.</p>
<p>Find the sum of all the multiples of $3$ or $5$ below $1000$.</p>
*/


func ListAllNaturalNumsBelowLimit(limit int) (int) {
  var sum int = 0
  for i := 1; i < limit; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  return sum
}





