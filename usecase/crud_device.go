package usecase

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "learning/models"
    "learning/repository"
    "learning/utils"
)

var reader = bufio.NewReader(os.Stdin)

func CreateDevice(lang string) {
    msg := utils.GetMessages(lang)

    fmt.Print(msg.PromptDeviceCode)
    code := readLine()

    fmt.Print(msg.PromptDeviceName)
    name := readLine()

    fmt.Print(msg.PromptLocation)
    location := readLine()

    fmt.Print(msg.PromptStatus)
    status := readLine()

    fmt.Print(msg.PromptDeviceType)
    deviceType := readLine()

    device := models.Device{
        DeviceCode: code,
        DeviceName: name,
        Location:   location,
        Status:     status,
        DeviceType: deviceType,
    }

    err := repository.CreateDevice(&device)
    if err != nil {
        fmt.Println(msg.ErrorOperation, err)
        return
    }

    fmt.Println(msg.CreateSuccess)
}

func ListDevices(lang string) {
    msg := utils.GetMessages(lang)

    devices, err := repository.GetAllDevices()
    if err != nil {
        fmt.Println(msg.ErrorOperation, err)
        return
    }
    if len(devices) == 0 {
        fmt.Println(msg.NoDevicesFound)
        return
    }


	fmt.Printf("%-4s %-15s %-20s %-15s %-10s %-15s\n",
    msg.HeaderID, msg.HeaderCode, msg.HeaderName, msg.HeaderLocation, msg.HeaderStatus, msg.HeaderDeviceType)
    for _, d := range devices {
        fmt.Printf("%-4d %-15s %-20s %-15s %-10s %-15s\n",
            d.ID, d.DeviceCode, d.DeviceName, d.Location, d.Status, d.DeviceType)
    }
}

func UpdateDevice(lang string) {
    msg := utils.GetMessages(lang)

    fmt.Print(msg.PromptDeviceID)
    id := readInt()

    device, err := repository.GetDeviceByID(id)
    if err != nil {
        fmt.Println(msg.ErrorOperation, err)
        return
    }

    fmt.Printf("%s (%s): ", msg.PromptDeviceCode, device.DeviceCode)
    device.DeviceCode = readDefault(device.DeviceCode)

    fmt.Printf("%s (%s): ", msg.PromptDeviceName, device.DeviceName)
    device.DeviceName = readDefault(device.DeviceName)

    fmt.Printf("%s (%s): ", msg.PromptLocation, device.Location)
    device.Location = readDefault(device.Location)

    fmt.Printf("%s (%s): ", msg.PromptStatus, device.Status)
    device.Status = readDefault(device.Status)

    fmt.Printf("%s (%s): ", msg.PromptDeviceType, device.DeviceType)
    device.DeviceType = readDefault(device.DeviceType)

    err = repository.UpdateDevice(&device)
    if err != nil {
        fmt.Println(msg.ErrorOperation, err)
        return
    }

    fmt.Println(msg.UpdateSuccess)
}

func DeleteDevice(lang string) {
    msg := utils.GetMessages(lang)

    fmt.Print(msg.PromptDeviceID)
    id := readInt()

    err := repository.DeleteDevice(id)
    if err != nil {
        fmt.Println(msg.ErrorOperation, err)
        return
    }

    fmt.Println(msg.DeleteSuccess)
}

func readLine() string {
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

func readDefault(defaultValue string) string {
    line := strings.TrimSpace(readLine())
    if line == "" {
        return defaultValue
    }
    return line
}

func readInt() int {
    for {
        text := readLine()
        n, err := strconv.Atoi(text)
        if err == nil {
            return n
        }
        fmt.Println("Input harus berupa angka.")
    }
}