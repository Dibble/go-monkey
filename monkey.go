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

func monkey(word string, result chan<- int) {
  guess := randString(len(word))
  guesses := 1

  for guess != word {
    guess = randString(len(word))
    guesses++
  }

  result <- guesses
}

func main() {
  rand.Seed(time.Now().UnixNano())

  word := randString(wordLength)
  fmt.Printf("Word to guess: %s\n", word)

  startTime := time.Now().UnixNano()

  result := make(chan int)
  go monkey(word, result)
  guesses := <- result

  endTime := time.Now().UnixNano()
  totalTime := endTime - startTime

  fmt.Println("Guessed the word!")
  fmt.Printf("In %d guesses.\n", guesses)
  fmt.Printf("Total time: %dns\n", totalTime)
  fmt.Printf("%dns per guess.\n", totalTime / int64(guesses))
}
