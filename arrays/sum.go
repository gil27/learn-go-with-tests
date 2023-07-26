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
