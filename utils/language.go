package utils

type LangData struct{
	MenuTitle string
	OptionAdd string
	OptionList string
	OptionUpdate string
	OptionDelete string
	OptionExit string
	PromptSelect string
	InvalidSelect string
}

var messages = map[string]LangData{
	"id": {
		MenuTitle:    "IOT MONITORING DEVICE CLI",
		OptionAdd:    "1. Buat Perangkat",
		OptionList:   "2. Daftar Perangkat",
		OptionUpdate: "3. Perbarui Perangkat",
		OptionDelete: "4. Hapus Perangkat",
		OptionExit:   "5. Keluar",
		PromptSelect: "Pilih opsi: ",
		InvalidSelect: "Opsi tidak valid. Silakan coba lagi.",
	},
	"en": {
		MenuTitle:    "IOT MONITORING DEVICE CLI",
		OptionAdd:    "1. Create Device",
		OptionList:   "2. List Devices",
		OptionUpdate: "3. Update Device",
		OptionDelete: "4. Delete Device",
		OptionExit:   "5. Exit",
		PromptSelect: "Select an option: ",
		InvalidSelect: "Invalid option. Please try again.",
	},
}

func GetMessages(lang string) LangData {
    if msg, ok := messages[lang]; ok {
        return msg
    }
    return messages["en"]
}

func GetAllLanguages() []string {
    return []string{"id", "en"}
}