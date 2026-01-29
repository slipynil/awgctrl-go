package awgctrlgo

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// fileName like "user" or "path/to/file/user", virtualEndpoint like "10.66.66.02/32"
// return filePath, peerPublicKey.String(), error
func (a *awg) AddPeer(fileName, virtualEndpoint string) (string, string, error) {

	// check endpoint format
	split := strings.Split(virtualEndpoint, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("invalid virtualEndpoint format")
	}
	// parse mask and IP virtual endpoint
	_, ipNet, err := net.ParseCIDR(virtualEndpoint)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse CIDR: %w", err)
	}

	// generate peer's private key
	peerPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %w", err)
	}

	// generate PresharedKey
	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate preshared key: %w", err)
	}

	// create configuration file for user
	filePath, err := a.createFileCfg(fileName, peerPrivateKey, presharedKey, virtualEndpoint)
	if err != nil {
		return "", "", fmt.Errorf("failed to create configuration file: %w", err)
	}

	peerPublicKey := peerPrivateKey.PublicKey()

	peerCfg := wgtypes.PeerConfig{
		PublicKey:    peerPublicKey,
		PresharedKey: &presharedKey,
		AllowedIPs:   []net.IPNet{*ipNet},
	}

	cfg := wgtypes.Config{
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{peerCfg},
	}

	// Set new device configuration (tunnel)
	if err := a.client.ConfigureDevice(a.deviceName, cfg); err != nil {
		return "", "", fmt.Errorf("failed to configure device: %w", err)
	}

	if a.debug {
		device, err := a.client.Device(a.deviceName)
		if err != nil {
			return "", "", fmt.Errorf("failed to get device: %w", err)
		}
		for _, peer := range device.Peers {
			if peer.PublicKey == peerPublicKey {
				peerInfo(peer)
				return filePath, peerPublicKey.String(), nil
			}
		}
	}

	return filePath, peerPublicKey.String(), nil
}
