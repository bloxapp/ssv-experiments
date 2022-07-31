package types

//go:generate go run .../fastssz/sszgen --path messages.go --include ./message_id.go --exclude-objs MessageID,MsgType
//go:generate go run .../fastssz/sszgen --path duty.go
//go:generate go run .../fastssz/sszgen --path consensus_input.go --include ./duty.go
