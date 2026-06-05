package usecase

import (
	"fmt"
	"learning/utils"
)

func ShowMenu(lang string) {
    msg := utils.GetMessages(lang)
    fmt.Println("=== " + msg.MenuTitle + " ===")
    fmt.Println(msg.OptionAdd)
    fmt.Println(msg.OptionList)
    fmt.Println(msg.OptionUpdate)
    fmt.Println(msg.OptionDelete)
    fmt.Println(msg.OptionExit)
    fmt.Print(msg.PromptSelect)
}
