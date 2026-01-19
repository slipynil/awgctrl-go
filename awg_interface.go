package awgctrlgo

import (
	"github.com/Jipok/wgctrl-go"
	"github.com/Jipok/wgctrl-go/wgtypes"
)

// interface for working with awg
type awgClient interface {
	ConfigureDevice(name string, cfg wgtypes.Config) error
	Device(name string) (*wgtypes.Device, error)
	Close() error
}

type awg struct {
	client      awgClient       // client for working with awg
	device      *wgtypes.Device // device for working with awg
	storagePath string          // path to create user.conf files
	endpoint    string          // IP:PORT
	obfuscation Obfuscation     // config for obfuscation
}

// Create new awg service,
// DOES NOT CREATE A NEW TUNNEL, BUT ONLY CONNECTS TO AN EXISTING TUNNEL
func New(tunnelName, endpoint, storagePath string, obfuscation *Obfuscation) (*awg, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	device, err := client.Device(tunnelName)
	if err != nil {
		return nil, err
	}
	return &awg{
		client:      client,
		device:      device,
		storagePath: storagePath,
		endpoint:    endpoint,
		obfuscation: *obfuscation,
	}, nil
}
