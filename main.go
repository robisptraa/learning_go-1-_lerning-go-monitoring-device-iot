package main

import (
	"fmt"
	"learning/database"
	"learning/repository"
	"learning/models"
)



func main() {
	database.Connect()

	for {
		fmt.Println("1. Create Device")
		fmt.Println("2. List Devices")
		fmt.Println("3. Update Device")
		fmt.Println("4. Delete Device")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")
	}
	fmt.Println("IOT MONITORING DEVICE CLI Mas OBets learning Go")

}

