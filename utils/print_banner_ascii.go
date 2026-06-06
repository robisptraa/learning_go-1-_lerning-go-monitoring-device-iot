package utils

import "fmt"


func PrintBannerWelcome() {
	fmt.Print()

 fmt.Println(
	`
	                                                    
__      _____| | ___ ___  _ __ ___   ___  | |_ ___            
\ \ /\ / / _ \ |/ __/ _ \| '_ _\ \ / _ \ | __/ _ \           
 \ V  V /  __/ | (_| (_) | | | | | |  __/ | || (_) |          
  \_/\_/ \___|_|\___\___/|_| |_| |_|\___|  \__\___/           
                                                              
       _          _                                   _       
  ___ | |__   ___| |_ ___    ___ ___  _ __  ___  ___ | | ___  
 / _ \| '_ \ / _ \ __/ __|  / __/ _ \| '_ \/ __|/ _ \| |/ _ \ 
| (_) | |_) |  __/ |_\__ \ | (_)| (_) | | | \__ \ (_) | |  __/ 
 \___/|_.__/ \___|\__|___/  \___\___/|_| |_|___/\___/|_|\___| 
                                                              
	`,
 )

 fmt.Println()
}

func PrintBannerMenu(lang string) {
    msg := GetMessages(lang)
    fmt.Println()            
    fmt.Print(msg.BannerAscii)   
    fmt.Println()          
}

