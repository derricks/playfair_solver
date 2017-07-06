package main

import (
  "fmt"
  "playfair_solver/wordsquare"
)

// This is specifically geared towards the Hills of Homicied problem
// in P&A July/August 2017
func main() {
   // args = plaintext digraph,ciphertext digraph, etc.
   // e.g.: playfair_solver FI,HX QZ,TV
   
   // for each line in stdin, create a word sqaure for it
   // then encrypt each digraph against word square and
   // see if it maps to the equivalent ciphertext. If all parse,
   // match.
   square := wordsquare.KeySquareFromString("monarchy")
   fmt.Println(square)
}