package main

import "fmt"
import "time"
import "math/rand"

var board [10][10]bool
var nextboard [10][10]bool

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  var t int
  t = rand.Intn(50)
  fmt.Println("Creating", t, "cells")
  for i := 0; i < t; i++ {
    board[rand.Intn(10)][rand.Intn(10)] = true
  }

  printboard()
  for ;; {
    tick()
    printboard()
    time.Sleep(time.Second)
  }
}

func tick() {
  for i, row := range(board) {
    for j, cell := range(row) {
      n := neighbours(i, j)
      if(cell) {
        if(n < 2 || n > 3) {
          nextboard[i][j] = false
        } else {
          nextboard[i][j] = true
        }
      } else {
        if(n == 3) {
          nextboard[i][j] = true
        } else {
          nextboard[i][j] = false
        }
      }
    }
  }

  for i, row := range(nextboard) {
    for j, cell := range(row) {
      board[i][j] = cell
    }
  }
}

func neighbours(x int, y int) int {
  count := 0
  for i := x - 1; i <= x + 1; i++ {
    for j := y - 1; j <= y + 1; j++ {
      if(0 <= i && i < 10 && 0 <= j && j < 10) {
        if(board[i][j]) {
          count++
        }
      }
    }
  }
  if(board[x][y]) {
    count--
  }
  return count
}

func printboard() {
  fmt.Print("\x0c")
  for _, row := range(board) {
    for _, cell := range(row) {
      if(cell) {
        fmt.Printf("X")
      } else {
        fmt.Printf(" ")
      }
    }
    fmt.Println()
  }
}
