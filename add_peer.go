package awgctrlgo

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// fileName like a "user" or "path/to/file/user",
// socket like a "10.66.66.02/32"
// return filePath, Peer struct, error
func (a *awg) AddPeer(fileName, socket string) (string, *Peer, error) {

	// check endpoint format
	split := strings.Split(socket, "/")
	if len(split) != 2 {
		return "", nil, fmt.Errorf("invalid socket format")
	}
	// parses mask and IP from socket
	_, ipNet, err := net.ParseCIDR(socket)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse CIDR: %w", err)
	}

	// generates peer's private key
	peerPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// generates peer's PresharedKey
	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate preshared key: %w", err)
	}

	// create configuration file for user
	filePath, err := a.createFileCfg(fileName, peerPrivateKey, presharedKey, socket)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create configuration file: %w", err)
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
		return "", nil, fmt.Errorf("failed to configure device: %w", err)
	}

	p := &Peer{
		PublicKey:     peerPublicKey.String(),
		PresharedKey:  presharedKey.String(),
		VirtualSocket: socket,
	}

	// if debug mode is enabled, print peer info
	if a.debug {
		p.Info()
	}

	return filePath, p, nil
}
