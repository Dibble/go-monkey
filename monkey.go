package main

import (
  "fmt"
  "math/rand"
  "time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"
func randString(length int) string {
  result := make([]byte, length)

  for i := range result {
      result[i] = letterBytes[rand.Intn(len(letterBytes))]
  }

  return string(result)
}

func printStats(monkey int, guesses int, totalTime time.Duration) {
  fmt.Printf("Monkey #%d: In %d guesses.\n", monkey, guesses)
  fmt.Printf("Total time: %.2f sseconds\n", totalTime.Seconds())
  fmt.Printf("%dns per guess.\n\n", totalTime.Nanoseconds() / int64(guesses))
}

func monkey(words <-chan string, result chan<- int) {
  current_word := ""
  guesses := 0

  for true {
    select {
    case word := <- words:
      guesses = 0
      current_word = word
    default:
      if current_word != "" {
        guess := randString(len(current_word))
        guesses++

        if guess == current_word {
          result <- guesses
          current_word = ""
        }
      }
    }
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())

  words_channels := make([]chan string, 4)
  result_channels := make([]chan int, 4)
  for j := 0; j < 4; j++ {
    words_channels[j] = make(chan string)
    result_channels[j] = make(chan int)

    go monkey(words_channels[j], result_channels[j])
  }

  for i := 4; i <= 8; i++ {
    word := randString(i)
    fmt.Printf("Word to guess: %s\n", word)

    startTime := time.Now()

    for _, word_channel := range words_channels {
      word_channel <- word
    }

    select {
    case guesses := <- result_channels[0]:
      totalTime := time.Since(startTime)
      printStats(0, guesses, totalTime)
    case guesses := <- result_channels[1]:
      totalTime := time.Since(startTime)
      printStats(1, guesses, totalTime)
    case guesses := <- result_channels[2]:
      totalTime := time.Since(startTime)
      printStats(2, guesses, totalTime)
    case guesses := <- result_channels[3]:
      totalTime := time.Since(startTime)
      printStats(3, guesses, totalTime)
    }
  }
}
