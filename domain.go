package awgctrlgo

// Obfuscation struct for initialization awg service
// use your values in /etc/amnezia/amneziawg/file.conf
type Obfuscation struct {
	Jc   int // Js counts junk packets (1 - 128)
	Jmin int // Minimum allowed Js value (Jmin < Jmax)
	Jmax int // Maximum allowd Js value (1 - 1280)
	S1   int // prefix init-packet
	S2   int // prefix response-packet
	H1   int // initiator to Responder (1 - 2**32-1)
	H2   int // Responder to initiator (1 - 2**32-1)
	H3   int // Data Packet (1 - 2**32-1)
	H4   int // Cookie Reply (1 - 2**32-1)
}
