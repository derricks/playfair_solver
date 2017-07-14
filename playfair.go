package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  
  "playfair_solver/wordsquare"
)

// This is specifically geared towards the Hills of Homicied problem
// in P&A July/August 2017.
// Usage: playfair_solver keyword_file plaintext_file ciphertext_to_match
func main() {
  keywordFileName := os.Args[1]
  plaintextFileName := os.Args[2]
  ciphertext := os.Args[3]
  
  cipherTextRegex, err := regexp.Compile(ciphertext)
  panicOnError(err)
  
  keywordFile, err := os.Open(keywordFileName)
  panicOnError(err)
  defer keywordFile.Close()
    
  keywordScanner := bufio.NewScanner(keywordFile)
  for keywordScanner.Scan() {
    keyword := keywordScanner.Text()
    checkKeywordAgainstPlaintext(keyword, plaintextFileName, cipherTextRegex)
  }     
}

func checkKeywordAgainstPlaintext(keyword string, plaintextFileName string, ciphertextRegex *regexp.Regexp) {
  square := wordsquare.KeySquareFromString(keyword)

  // now go through each plain text guess and compare encoded version
  // to ciphertext
  plaintextFile, err := os.Open(plaintextFileName)
  defer plaintextFile.Close()
  panicOnError(err)
  plaintextScanner := bufio.NewScanner(plaintextFile)
  for plaintextScanner.Scan() {
    plaintext := plaintextScanner.Text()
    encrypted := square.EncryptString(plaintext)
    
    if ciphertextRegex.Match([]byte(encrypted)) {
      fmt.Printf("keyword: %s, plaintext: %s, encrypted: %s\n", keyword, plaintext, encrypted)
    }
    
  }
}

func panicOnError(e error) {
  if e != nil {
      panic(e)
  }
}