package qbft

//go:generate rm -f ./types_encoding.go
//go:generate rm -f ./messages_encoding.go
//go:generate go run .../fastssz/sszgen --path . --include ../types --exclude-objs MsgContainer,FutureMsgContainer
