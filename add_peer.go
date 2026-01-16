package awgctrlgo

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// fileName like "user" or "path/to/file/user", virtualEndpoint like "10.66.66.02/32"
func (a *awg) AddPeer(fileName, virtualEndpoint string) (string, error) {

	// check endpoint format
	split := strings.Split(virtualEndpoint, "/")
	if len(split) != 2 {
		return "", fmt.Errorf("invalid virtualEndpoint format")
	}
	// get ip part
	ip := split[1]

	// check if the ip address is available
	if !a.isAllowedIP(ip) {
		return "", fmt.Errorf("no available IP")
	}
	// parse mask and IP virtual endpoint
	_, ipNet, err := net.ParseCIDR(virtualEndpoint)
	if err != nil {
		return "", err
	}

	// generate peer's private key
	peerPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
	}

	// generate PresharedKey
	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("failed to generate preshared key: %w", err)
	}

	// create configuration file for user
	if err := a.createFileCfg(
		fileName,
		peerPrivateKey,
		presharedKey,
		virtualEndpoint,
	); err != nil {
		return "", err
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
	if err := a.client.ConfigureDevice(a.device.Name, cfg); err != nil {
		return "", err
	}

	return peerPublicKey.String(), nil
}

// check IP is available
func (a *awg) isAllowedIP(ip string) bool {

	// go through all peers
	for _, peer := range a.device.Peers {
		// go through all occupied IPs of the peer
		for _, usedIP := range peer.AllowedIPs {
			// check if IP is already used
			if ip == usedIP.IP.String() {
				return false
			}
		}
	}
	return true
}
