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

func (square KeySquare) EncryptString(plaintext string) string {
  var returnBuffer bytes.Buffer
  
  // todo: strip non-letters
  dedupedString := padOutDoubleLetters(strings.ToUpper(plaintext))
  if len(dedupedString) % 2 == 1 {
    dedupedString = dedupedString + "X"
  }
  
  for digraphStart := 0; digraphStart < len(dedupedString) - 1; digraphStart += 2 {
   digraph := []string{string(dedupedString[digraphStart]), string(dedupedString[digraphStart + 1])}
   returnBuffer.WriteString(square.encodeDigraph(digraph))   
  }
  return returnBuffer.String()
}

func (square KeySquare) encodeDigraph(digraph []string) string {  
  first_coords := square.index[digraph[0]]
  second_coords := square.index[digraph[1]]
  
  if (first_coords.row != second_coords.row && first_coords.column != second_coords.column) {
    return square.Get(first_coords.row, second_coords.column) + 
           square.Get(second_coords.row, first_coords.column)
  } else if first_coords.row == second_coords.row {
    // wraparound case
    if second_coords.column == 4 {
      return square.Get(first_coords.row, first_coords.column + 1) +
             square.Get(second_coords.row, 0)
    } else if first_coords.column == 4 {
      return square.Get(first_coords.row, 0) +
             square.Get(second_coords.row, second_coords.column+1)
      
    } else {
      return square.Get(first_coords.row, first_coords.column + 1) +
             square.Get(second_coords.row, second_coords.column + 1)
    }
  } else {
    // columns are equal
    if second_coords.row == 4 {
      return square.Get(first_coords.row + 1, first_coords.column) +
             square.Get(0, second_coords.column)
    } else if first_coords.row == 4 {
      return square.Get(0, first_coords.column) +
             square.Get(second_coords.row + 1, second_coords.column)
    } else {
      return square.Get(first_coords.row + 1, first_coords.column) +
             square.Get(second_coords.row + 1, second_coords.column)
    }    
  }  
}

// playfair specifies that any doubled letters should be interrupted with X's
// e.g.: HILLSTREETBLUES becomes HILXLSTREXETBLUES
func padOutDoubleLetters(instring string) string {
  var buffer bytes.Buffer
  previousLetter := ""
  letters := strings.Split(instring, "")
  for _, letter := range letters {
    if letter == previousLetter {
      buffer.WriteString("X")
    }
    
    buffer.WriteString(letter)
    previousLetter = letter
  }
  return buffer.String()
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
