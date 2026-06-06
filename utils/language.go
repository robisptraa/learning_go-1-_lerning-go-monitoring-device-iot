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
	PromptDeviceCode string
    PromptDeviceName string
    PromptLocation   string
    PromptStatus     string 
    PromptDeviceID   string
	ErrorOperation  string
    CreateSuccess   string
	NoDevicesFound  string
	UpdateSuccess   string
	DeleteSuccess   string
}

var messages = map[string]LangData{
	"id": {
		MenuTitle:    "OBETS MONITORING DEVICE IOT BASE CLI",
		OptionAdd:    "1. Buat Perangkat",
		OptionList:   "2. Daftar Perangkat",
		OptionUpdate: "3. Perbarui Perangkat",
		OptionDelete: "4. Hapus Perangkat",
		OptionExit:   "5. Keluar",
		PromptSelect: "Pilih opsi: ",
		InvalidSelect: "Opsi tidak valid. Silakan coba lagi.",
		PromptDeviceCode: "Masukkan kode perangkat: ",
    	PromptDeviceName: "Masukkan nama perangkat: ",
    	PromptLocation:   "Masukkan lokasi: ",
    	PromptStatus:     "Masukkan status perangkat: ", 
    	PromptDeviceID:   "Masukkan ID perangkat: ",
		ErrorOperation:  "Terjadi kesalahan: ",
		CreateSuccess:   "Perangkat berhasil dibuat.",
		NoDevicesFound:  "Perangkat tidak ditemukan.",
		UpdateSuccess:   "Perangkat berhasil diperbarui.",
		DeleteSuccess:   "Perangkat berhasil dihapus.",
	},
	"en": {
		MenuTitle:    "OBETS MONITORING DEVICE IOT BASE CLI",
		OptionAdd:    "1. Create Device",
		OptionList:   "2. List Devices",
		OptionUpdate: "3. Update Device",
		OptionDelete: "4. Delete Device",
		OptionExit:   "5. Exit",
		PromptSelect: "Select an option: ",
		InvalidSelect: "Invalid option. Please try again.",
		PromptDeviceCode: "Enter device code: ",
		PromptDeviceName: "Enter device name: ",
		PromptLocation:   "Enter location: ",
		PromptStatus:     "Enter device status: ",
		PromptDeviceID:   "Enter device ID: ",
		ErrorOperation:  "An error occurred: ",
		CreateSuccess:   "Device created successfully.",
		NoDevicesFound:  "No devices found.",
		UpdateSuccess:   "Device updated successfully.",
		DeleteSuccess:   "Device deleted successfully.",
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