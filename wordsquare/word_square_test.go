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

func TestPadDoubleLettersWithX(test *testing.T) {
  paddedString := padOutDoubleLetters("AFFLUX")
  if paddedString != "AFXFLUX" {
    test.Fatalf("Expected AFXFLUX, got %s", paddedString)
  }
}

func TestEncryption(test *testing.T) {
 square := KeySquareFromString("monarchy")
 encodedString := square.EncryptString("wearediscoveredsaveyourself")
 if encodedString != "UGRMKCSXHMUFMKBTOXGCMVATLUIV" {
   test.Fatalf("Expected UGRMKCSXHMUFMKBTOXGCMVATLUIV, but got %s", encodedString)
 }
}

func TestPandaExamples(test *testing.T) {
  testCases := [][]string {
    {"NYPDBLUE", "HILLSTREETBLUES", "IKATCMXMAWLWNCEARZ"},
    {"BROOKLYNNINENINE", "HOMICIDELIFEONTHESTREXET", "FLQYFYGNOAGIRIZTGXPLGKAS"},
    {"THESHIELD", "NCISNEWORLEANS", "UNTIKIVPOBAKMI"},
    {"HOMICIDELIFEONTHESTREET", "LAWORDERSPECIALVICTIMSUNIT", "FROETLLSRKNOFQDXCHAHORZBHA"},
    {"CSIMIAMI", "NYPDBLUE", "TMQBEHRG"},
    {"CRIMINALMINDS", "CSIMIAMI", "MAMNCDNM"},
  }
  
  for _, testCase := range testCases {
    square := KeySquareFromString(testCase[0])
    encodedString := square.EncryptString(testCase[1])
    
    if encodedString != testCase[2] {
      test.Fatalf("With keyword %s, expected %s to encode to %s. Was: %s", testCase[0], testCase[1], testCase[2], encodedString)
    }
  }

}

