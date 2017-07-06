package wordsquare

import (
  "bytes"
  "strings"
)
type coord struct {
  row, column int
}

type KeySquare struct {
  square [5][5]string
  index  map[string]coord
  insert_row, insert_column int
}

var alphabet = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
                            "S", "T", "U", "V", "W", "X", "Y", "Z"}
                            
func KeySquareFromString(keyword string) KeySquare {
  normalized_keyword := strings.TrimSpace(strings.ToUpper(keyword))
  all_letters := strings.Split(normalized_keyword, "")
  returnSquare := KeySquare{index:make(map[string]coord)}
  returnSquare.addLetters(all_letters)
  // since alphabet is an array, we have to convert it to a slice
  // to make the types line up
  returnSquare.addLetters(alphabet[:])
  
  return returnSquare
}

func (square KeySquare) Get(row int, column int) string {
  return square.square[column][row]
}

func (square KeySquare) EncryptString(digraph string) string {
  //letters := strings.Split(digraph, "")
  
  return ""
}

func (square KeySquare) String() string {
  var buffer bytes.Buffer
  for row := 0; row < 5; row++ {
    for column := 0; column < 5; column++ {
      buffer.WriteString(square.Get(row, column))
      buffer.WriteString(" ")
    }
    buffer.WriteString("\n")
  }
  return buffer.String()
}

// private methods
func (square *KeySquare) addLetters(letters []string) {
  for _, letter := range letters {
    if (letter == "J") {
      square.addLetter("I")  
    } else {
      square.addLetter(letter)
    }
  }   
}

// add the letter to the square and update all the data structures
func (square *KeySquare) addLetter(letter string) {
  if _, found := square.index[letter]; !found {
    square.square[square.insert_column][square.insert_row] = letter
    square.index[letter] = coord{row: square.insert_row, column: square.insert_column}
    
    square.insert_column = square.insert_column + 1
    if square.insert_column >= 5 {
      square.insert_row = square.insert_row + 1
      square.insert_column = 0
    }
  }
  
}
