package awgctrlgo

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// RestorePeer registers an existing peer back into WireGuard using its known keys.
// Does NOT generate new keys and does NOT create a new config file.
// Used when renewing a subscription — the user's existing .conf file remains valid.
func (a *awg) RestorePeer(publicKeyStr, presharedKeyStr, socket string) error {
	split := strings.Split(socket, "/")
	if len(split) != 2 {
		return fmt.Errorf("invalid socket format")
	}

	_, ipNet, err := net.ParseCIDR(socket)
	if err != nil {
		return fmt.Errorf("failed to parse CIDR: %w", err)
	}

	publicKey, err := wgtypes.ParseKey(publicKeyStr)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %w", err)
	}

	presharedKey, err := wgtypes.ParseKey(presharedKeyStr)
	if err != nil {
		return fmt.Errorf("failed to parse preshared key: %w", err)
	}

	peerCfg := wgtypes.PeerConfig{
		PublicKey:    publicKey,
		PresharedKey: &presharedKey,
		AllowedIPs:   []net.IPNet{*ipNet},
	}

	cfg := wgtypes.Config{
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{peerCfg},
	}

	return a.client.ConfigureDevice(a.deviceName, cfg)
}
