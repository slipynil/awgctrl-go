# awgctrl-go

## What is this?
This is a Go package that provides a simple interface to manage amneziawg devices.

## How to use?
### Installation
To install the package, run the following command:
```
go get github.com/slipynil/awgctrl-go
```

### Code example

```go
package main

import (
	"os"
	"time"
	"path/filepath"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)
func main() {
	// USE values ONLY ACTIVE tunnel's CONFIGURATION
	// also in /etc/amnezia/amneziawg/awg0.conf
	// available only this fields
	cfg := awgctrlgo.Obfuscation{
		Jc: 2,
		Jmin: 10,
		Jmax: 50,
		S1: 10,
		S2: 20,
		H1: 123456,
		H2: 234567,
		H3: 345678,
		H4: 456789,
	}

	tunnelName := "awg0"
	endpoint := "localhost:5050"

	// client for managing amneziawg devices
	// Not creating a new tunnel, using existing one
	storagePath, _ := filepath.Abs("./data")
	awg, err := awgctrlgo.New(tunnelName, endpoint, storagePath, &cfg)
	if err != nil {
		panic(err)
	}
	defer awg.Close()

	// information about the tunnel
	awg.DeviceInfo()

	// create a new peer
	userPublicKey, err := awg.AddPeer("user", "10.66.66.02/32")
	if err != nil {
		panic(err)
	}
	// information about the peers
	awg.ShowPeers()

	// delete a peer by public key
	time.Sleep(time.Minute * 5)
	if err := awg.DeletePeer(userPublicKey); err != nil {
		panic(err)
	}
}
```
