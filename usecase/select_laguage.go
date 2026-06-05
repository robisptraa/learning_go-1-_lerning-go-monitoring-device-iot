package usecase

import (
	"fmt"
)

func SelectLanguage() string {
    fmt.Println("Choose language / Pilih bahasa")
    fmt.Println("1. English")
    fmt.Println("2. Bahasa Indonesia")
    fmt.Print("Choice: ")

    var langChoice int
    fmt.Scanln(&langChoice)

    if langChoice == 2 {
        return "id"
    }
    return "en"
}
