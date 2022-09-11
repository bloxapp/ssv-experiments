package ssv

//go:generate rm -f ./messages_encoding.go
//go:generate rm -f ./runner_state_encoding.go
//go:generate go run .../fastssz/sszgen --path . --include ../qbft,../types --exclude-objs PartialSigContainer
