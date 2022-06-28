package main

import "fmt"

func main() {
  var roman string
  decimal := 0
  
  // Taking input
  fmt.Print("Enter you Roman numeral: ")
  fmt.Scan(&roman)

  for i := 0 ; i < len(roman) ; i++ {
    if roman[i] == 'I' { decimal += 1 } else
    if roman[i] == 'V' { decimal += 5 } else
    if roman[i] == 'X' { decimal += 10 } else
    if roman[i] == 'L' { decimal += 50 } else
    if roman[i] == 'C' { decimal += 100 } else
    if roman[i] == 'D' { decimal += 500 } else
    if roman[i] == 'M' { decimal += 1000 }

    // Dealing edge cases like IV = 4, IX = 9 and so on
    if (i+1) < len(roman) {
      if roman[i] == 'I' && roman[i+1] == 'V' { decimal -= 2} else
      if roman[i] == 'I' && roman[i+1] == 'X' { decimal -= 2} else
      if roman[i] == 'X' && roman[i+1] == 'L' { decimal -= 20} else
      if roman[i] == 'X' && roman[i+1] == 'C' { decimal -= 20} else
      if roman[i] == 'C' && roman[i+1] == 'D' { decimal -= 200} else
      if roman[i] == 'C' && roman[i+1] == 'M' { decimal -= 200}
    }
  }
  fmt.Println(decimal)
}
