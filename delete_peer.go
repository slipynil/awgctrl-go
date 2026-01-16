package awgctrlgo

import "github.com/Jipok/wgctrl-go/wgtypes"

// DeletePeer deletes a peer from the tunnel by peer's public key
func (a *awg) DeletePeer(peerPublicKeyStr string) error {
	peerPublicKey, err := wgtypes.ParseKey(peerPublicKeyStr)
	if err != nil {
		return err
	}
	cfg := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			// peerConfig - the configuration for a peer. In it, you specify what to do with a specific peer
			{
				PublicKey: peerPublicKey, // public key of the peer
				Remove:    true,          // true = delete this peer
			},
		},
	}
	// apply the configuration to the device (tunnel)
	return a.client.ConfigureDevice(a.device.Name, cfg)
}
