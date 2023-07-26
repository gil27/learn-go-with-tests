package main

func Sum(numbers []int) (total int) {
  // for i := 0; i < 5; i++ {
  //   total += numbers[i]
  // }

  for _, number := range numbers {
    total += number
  }
  return
}

func SumAll(slices ...[]int) []int {
  // slicesLen := len(slices)
  // total := make([]int, slicesLen)
  //
  // for i, slice := range slices {
  //   total[i] = Sum(slice)
  // }


  var total []int

  for _, slice := range slices {
    total = append(total, Sum(slice))
  }

  return total
}

func SumAllTails(numbersToSum ...[]int) []int {
  var total []int

  for _, numbers := range numbersToSum {
    if len(numbers) == 0 {
      total = append(total, 0)
    } else {
      numbersToSum := numbers[1:]

      total = append(total, Sum(numbersToSum))
    }
  }

  return total
}
