package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishPrefix = "Hello, "
const spanishPrefix = "Holla, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }

  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
    prefix = frenchPrefix
  case spanish:
    prefix = spanishPrefix
  default:
    prefix = englishPrefix
  }

  return
}

func main() {
  fmt.Println(Hello("Gil", ""))
}
