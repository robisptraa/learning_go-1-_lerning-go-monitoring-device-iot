package repository

import (
	"learning/database"
	"learning/models"
)

func CreateDevice(device *models.Device) error {
	query := `INSERT INTO devices (
		device_code,
		device_name,
		location,
		status
	) VALUES (?, ?, ?, ?)`

	_, err := database.DB.Exec(
		query,
		device.DeviceCode,
		device.DeviceName,
		device.Location,
		device.Status,
	)

	return err
}

func GetAllDevices() ([]models.Device, error) {
	rows, err := database.DB.Query(`SELECT id, device_code, device_name, location, status FROM devices`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.DeviceCode, &device.DeviceName, &device.Location, &device.Status); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func GetDeviceByID(id int) (models.Device, error) {
    var device models.Device
    err := database.DB.QueryRow(
        `SELECT id, device_code, device_name, location, status FROM devices WHERE id = ?`,
        id,
    ).Scan(
        &device.ID,
        &device.DeviceCode,
        &device.DeviceName,
        &device.Location,
        &device.Status,
    )
    return device, err
}

func UpdateDevice(device *models.Device) error {
	query := `UPDATE devices SET device_code = ?, device_name = ?, location = ?, status = ? WHERE id = ?`

	_, err := database.DB.Exec(
		query,
		device.DeviceCode,
		device.DeviceName,
		device.Location,
		device.Status,
		device.ID,
	)

	return err
}

func DeleteDevice(id int) error {
	query := `DELETE FROM devices WHERE id = ?`

	_, err := database.DB.Exec(query, id)

	return err
}

