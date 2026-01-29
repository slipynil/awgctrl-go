package awgctrlgo

import (
	"fmt"
)

// DeviceInfo prints information about the device (tunnel)
func (a *awg) DeviceInfo() error {
	device, err := a.client.Device(a.deviceName)
	if err != nil {
		return fmt.Errorf("failed to get device info: %w", err)
	}
	fmt.Println("----amneziawg is running----")
	fmt.Println("Interface:", device.Name)
	fmt.Println("Private key:", device.PrivateKey)
	fmt.Println("Public key:", device.PublicKey)
	fmt.Println("Listen Port:", device.ListenPort)
	fmt.Println("Is amnezia:", device.IsAmnezia)

	return nil
}
