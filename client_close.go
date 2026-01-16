package awgctrlgo

// Close closes a connection to the tunnel.
func (a *awg) Close() error {
	return a.client.Close()
}
