package types

//go:generate rm -f ./consensus_input_encoding.go
//go:generate rm -f ./duty_encoding.go
//go:generate rm -f ./messages_encoding.go
//go:generate rm -f ./share_encoding.go

//go:generate go run .../fastssz/sszgen --path messages.go --include ./message_id.go --exclude-objs MessageID,MsgType
//go:generate go run .../fastssz/sszgen --path duty.go
//go:generate go run .../fastssz/sszgen --path consensus_input.go --include ./duty.go
//go:generate go run .../fastssz/sszgen --path share.go --include ./crypto.go --exclude-objs DomainType
