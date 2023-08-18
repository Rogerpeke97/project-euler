package problem5


/*
<p>$2520$ is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.</p>
<p>What is the smallest positive number that is <dfn class="tooltip">evenly divisible<span class="tooltiptext">divisible with no remainder</span></dfn> by all of the numbers from $1$ to $20$?</p>
*/

func SmallestPositiveNumberEvenNoRemainder () (int){
  for i := 10; i == i; i++ {
    for j := 2; j < 21; j++ {
      if i % j != 0 {
        break
      }
      if j == 20 {
        return i
      }
    }
  }
  return 0
}

