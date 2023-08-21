package problem7

/*
<p>By listing the first six prime numbers: $2, 3, 5, 7, 11$, and $13$, we can see that the $6$th prime is $13$.</p>
<p>What is the $10\,001$st prime number?</p>
*/

func Prime10001st () (int){
  var count int
  for i := 2; i == i; i++ {
    for j := 1; j < i; j++ {
      if j != 1 && i % j == 0 {
        break
      }
      if j == i - 1 {
        count++
        if count == 10001 {
          return i
        }
      }
    }
  }
  return 0
}
