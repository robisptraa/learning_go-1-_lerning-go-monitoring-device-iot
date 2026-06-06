package main

import (
	"fmt"
	"learning/database"
	// "learning/repository"
	// "learning/models"
	"learning/usecase"
	"learning/utils"
	
)

func main() {
	utils.PrintBannerWelcome()
	database.Connect()
	lang := usecase.SelectLanguage()
	_ = utils.GetMessages(lang)
	utils.PrintBannerMenu(lang)

	 fmt.Println(utils.GetMessages(lang).MenuTitle)

	for {
		usecase.ShowMenu(lang)
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			usecase.CreateDevice(lang)
		case 2:
			usecase.ListDevices(lang)
		case 3:
			usecase.UpdateDevice(lang)
		case 4:
			usecase.DeleteDevice(lang)
		case 5:
			fmt.Println(utils.GetMessages(lang).ExitMessage)
			return
		default:
			fmt.Println(utils.GetMessages(lang).InvalidSelect)
		}
	   
	}
}

