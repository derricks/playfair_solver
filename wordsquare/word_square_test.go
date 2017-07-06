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

func TestEncryption(test *testing.T) {
 square := KeySquareFromString("monarchy")
 encodedString := square.EncryptString("wearediscoveredsaveyourself")
 if encodedString != "UGRMKCSXHMUFMKBTOXGCMVATLUIV" {
   test.Fatalf("Expected UGRMKCSXHMUFMKBTOXGCMVATLUIV, but got %s", encodedString)
 }
}