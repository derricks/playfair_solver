package wordsquare

import (
  "testing"
)

func TestMonarchyWordSquare(test *testing.T) {
  square := KeySquareFromString("monarchy")
  
  if (square.Get(2,1) != "F") {
    test.Fatalf("coord 2,1 should be F, was %s", square.Get(2,1))
  }
}

func TestPadDoubleLetters(test *testing.T) {
  paddedString := padOutDoubleLetters("HILLSTREETBLUES")
  if (paddedString != "HILXLSTREXETBLUES") {
    test.Fatalf("Expected HILXLSTREXETBLUES, got %s", paddedString)
  }
}