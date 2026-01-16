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
	endpoint    string          // IP:PORT
	obfuscation Obfuscation     // config for obfuscation
	client      awgClient       // client for working with awg
	device      *wgtypes.Device // device for working with awg
}

// Create new awg service,
// DOES NOT CREATE A NEW TUNNEL, BUT ONLY CONNECTS TO AN EXISTING TUNNEL
func New(tunnelName string, endpoint string, obfuscation *Obfuscation) (*awg, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	device, err := client.Device(tunnelName)
	if err != nil {
		return nil, err
	}
	return &awg{
		endpoint:    endpoint,
		obfuscation: *obfuscation,
		client:      client,
		device:      device,
	}, nil
}
