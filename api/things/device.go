package things

type Device struct {
	ID                      string
	ProductID               string
	ProductSpecificDeviceID string
}

type UserDevice struct {
	UserID            string
	Device            Device
	DeviceDisplayName string
}

type UserDevices struct {
	Items []UserDevice
}
