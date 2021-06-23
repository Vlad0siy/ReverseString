package main

import (
  "fmt"
)

func ReverseStr(s string) string {
  runes := []rune(s)
  length := len(runes)
  for i := 0; i < length/2; i++ {
    j := length - i - 1
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func main() {
  text := ""
  fmt.Println("Enter string:")
  fmt.Scan(&text)
  fmt.Println(ReverseStr(text))
}
