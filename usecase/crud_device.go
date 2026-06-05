package usecase

import (
	"fmt"
	"learning/utils"
)


func CreateDevice(lang string) {
    msg := utils.GetMessages(lang)
    fmt.Println(msg.OptionAdd)
}

func ListDevices(lang string) {
    msg := utils.GetMessages(lang)
    fmt.Println(msg.OptionList)
}

func UpdateDevice(lang string) {
    msg := utils.GetMessages(lang)
    fmt.Println(msg.OptionUpdate)
}

func DeleteDevice(lang string) {
    msg := utils.GetMessages(lang)
    fmt.Println(msg.OptionDelete)
}