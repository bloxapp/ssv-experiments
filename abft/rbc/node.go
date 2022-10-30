package rbc

type Node interface {
	// Broadcast returns the reliably broadcasted value, nil if not
	Broadcast(leader int, data []byte) []byte
}
