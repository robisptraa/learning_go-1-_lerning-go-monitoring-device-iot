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
	database.Connect()
	lang := usecase.SelectLanguage()
	_ = utils.GetMessages(lang)

	 fmt.Println(utils.GetMessages(lang).MenuTitle)

	for {
		usecase.ShowMenu(utils.GetMessages(lang).MenuTitle)
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			usecase.CreateDevice(utils.GetMessages(lang).MenuTitle)
		case 2:
			usecase.ListDevices(utils.GetMessages(lang).MenuTitle)
		case 3:
			usecase.UpdateDevice(utils.GetMessages(lang).MenuTitle)
		case 4:
			usecase.DeleteDevice(utils.GetMessages(lang).MenuTitle)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println(utils.GetMessages(lang).InvalidSelect)
		}
	   
	}
}

