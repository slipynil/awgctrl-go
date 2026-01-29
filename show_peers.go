package awgctrlgo

import (
	"fmt"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// ShowPeers prints information about the connected peers
func (a *awg) ShowPeers() error {
	device, err := a.client.Device(a.deviceName)
	if err != nil {
		return fmt.Errorf("Failed to retrieve device information: %w", err)
	}
	if len(device.Peers) == 0 {
		fmt.Println("No peers connected")
	}

	for _, peer := range device.Peers {
		peerInfo(peer)
	}
	return nil
}

func peerInfo(peer wgtypes.Peer) {
	fmt.Println("---PEER CONNECTION---")
	fmt.Println("Public Key:", peer.PublicKey)
	fmt.Print("Endpoint IP:", peer.Endpoint.IP)
	fmt.Print("Endpoint port:", peer.Endpoint.Port)
	fmt.Println("Endpoint zone:", peer.Endpoint.Zone)
	fmt.Println("Last Handshake:", peer.LastHandshakeTime)
}
