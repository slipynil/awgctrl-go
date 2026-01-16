package awgctrlgo

import (
	"fmt"
)

// DeviceInfo prints information about the device (tunnel)
func (a *awg) DeviceInfo() {

	fmt.Println("----amneziawg is running----")
	fmt.Println("Interface:", a.device.Name)
	fmt.Println("Private key:", a.device.PrivateKey)
	fmt.Println("Public key:", a.device.PublicKey)
	fmt.Println("Listen Port:", a.device.ListenPort)
	fmt.Println("Is amnezia:", a.device.IsAmnezia)

}
