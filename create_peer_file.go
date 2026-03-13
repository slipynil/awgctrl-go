package awgctrlgo

import (
	"fmt"
	"os"
	"path/filepath"
)

// creates a new configuration file for user connection to the tunnel
func (a *awg) createFileCfg(
	fileName string,
	dnsFormat string,
	peer *Peer,
) (string, error) {
	device, err := a.client.Device(a.deviceName)
	if err != nil {
		return "", err
	}
	publicDeviceKey := device.PublicKey.String()

	var dns string
	if dnsFormat == "" {
		dns = "none"
	} else {
		dns = dnsFormat
	}

	str := fmt.Sprintf(`
[Interface]
PrivateKey = %s
Address = %s
DNS = %s
Jc = %v
Jmin = %v
Jmax = %v
S1 = %v
S2 = %v
H1 = %v
H2 = %v
H3 = %v
H4 = %v

[Peer]
PublicKey = %v
PresharedKey = %v
Endpoint = %v
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`,
		peer.PrivateKey,
		peer.VirtualSocket,
		dns,
		a.obfuscation.Jc,
		a.obfuscation.Jmin,
		a.obfuscation.Jmax,
		a.obfuscation.S1,
		a.obfuscation.S2,
		a.obfuscation.H1,
		a.obfuscation.H2,
		a.obfuscation.H3,
		a.obfuscation.H4,
		publicDeviceKey,
		peer.PresharedKey,
		a.endpoint,
	)

	// create configuration file for user
	filePath := filepath.Join(a.storagePath, fileName+".conf")
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file with filePath %v: %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.Write([]byte(str)); err != nil {
		return "", fmt.Errorf("failed to write to file in path %v: %w", filePath, err)
	}

	return filePath, nil
}
