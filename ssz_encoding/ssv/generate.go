package ssv

//go:generate rm -f ./messages_encoding.go
//go:generate rm -f ./state_encoding.go
//go:generate rm -f ./base_runner_encoding.go

//go:generate go run .../fastssz/sszgen --path . --include ../qbft,../types --exclude-objs PartialSigContainer
