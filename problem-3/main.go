package problem3

/*
<p>The prime factors of 13195 are 5, 7, 13 and 29.</p>
<p>What is the largest prime factor of the number 600851475143?</p>
*/


func FindLargestPrimeFactor () (int) {
	var num_to_find_largest_prime = 600851475143
  var largest_prime_factor int
  var num_to_factorize = num_to_find_largest_prime
	for i := 1; i < num_to_find_largest_prime; i++ {
    if num_to_factorize == 1 {
      break
    }
    for j := 1; j < i; j++ {
      if i % j != 0 && j != 1 && j != i && num_to_factorize % i == 0 {
        largest_prime_factor = i
        num_to_factorize = num_to_factorize / i
      }
    }
	}
  return largest_prime_factor
}

