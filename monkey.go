package main

import (
  "fmt"
  "math/rand"
  "time"
)

const wordLength = 4
const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func randString(length int) string {
  result := make([]byte, length)

  for i := range result {
      result[i] = letterBytes[rand.Intn(len(letterBytes))]
  }

  return string(result)
}

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Monkeys!")

  word := randString(wordLength)
  fmt.Printf("Word to guess: %s\n", word)

  startTime := time.Now().UnixNano()
  guess := randString(wordLength)
  guesses := 1

  for guess != word {
    guess = randString(wordLength)
    guesses++
  }
  endTime := time.Now().UnixNano()

  totalTime := endTime - startTime

  fmt.Println("Guessed the word!")
  fmt.Printf("In %d iterations.\n", guesses)
  fmt.Printf("Total time: %dns\n", totalTime)
  fmt.Printf("%dns per guess.\n", totalTime / int64(guesses))
}
