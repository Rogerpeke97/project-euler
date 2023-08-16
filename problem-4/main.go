package problem4


func isPalindrome (number int) (bool) {
  var num_to_modify = number
  var reversed_number int
  for num_to_modify > 0 {
    reversed_number = reversed_number * 10 + num_to_modify % 10
    num_to_modify /= 10
  }
  return reversed_number == number
}

func LargestPalindromeProductOfThreeDigits() (int) {
  var highest_palindrome int
  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {
      if isPalindrome(i * j) && i * j > highest_palindrome {
        highest_palindrome = i * j        
      }
    }
  }
  return highest_palindrome
}
