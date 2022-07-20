package qbft

//go:generate go run .../fastssz/sszgen --path . [--objs Message,SignedMessage]

type Message struct {
}

type SignedMessage struct {
	Message *Message
}
