package dkg

type Message struct {
	DataSSZSnappy []byte `ssz-max:"2048"`
}

type SignedMessage struct {
	Message   Message
	Signer    uint64
	Signature [65]byte `ssz-size:"65"`
}

// Init is the first message in a DKG which initiates a DKG
type Init struct {
	// OperatorIDs are the operators selected for the DKG
	OperatorIDs []uint64 `ssz-max:"13"`
	// Threshold DKG threshold for signature reconstruction
	Threshold uint16
	// WithdrawalCredentials used when signing the deposit data
	WithdrawalCredentials []byte `ssz-size:"32"`
	// Fork is eth2 fork version
	Fork [4]byte `ssz-size:"4"`
}

// Output is the last message in every DKG which marks a specific node's end of process
type Output struct {
	// EncryptedShare standard SSV encrypted shares
	EncryptedShare [256]byte `ssz-size:"256"`
	// SharePubKey is the share's BLS pubkey
	SharePubKey [48]byte `ssz-size:"48"`
	// ValidatorPubKey the resulting public key corresponding to the shared private key
	ValidatorPubKey [48]byte `ssz-size:"48"`
	// DepositDataSignature reconstructed signature of DepositMessage according to eth2 spec
	DepositDataSignature [96]byte `ssz-size:"96"`
}

// SignedOutput signs the Output message with an EVM compatible signature
type SignedOutput struct {
	// Data signed
	Data Output
	// Signer Operator ID which signed
	Signer uint64
	// Signature ECDSA signature over an Output msg
	Signature [65]byte `ssz-size:"65"`
}
