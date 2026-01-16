package awgctrlgo

import "fmt"

// ShowPeers prints information about the connected peers
func (a *awg) ShowPeers() {
	if len(a.device.Peers) == 0 {
		fmt.Println("No peers connected")
	}

	for _, peer := range a.device.Peers {
		fmt.Println("---PEER CONNECTION---")
		fmt.Println("Public Key:", peer.PublicKey)
		fmt.Println("Last Handshake:", peer.LastHandshakeTime)
	}
}
