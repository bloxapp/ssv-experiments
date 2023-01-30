package types

//go:generate rm -f ./consensus_data_encoding.go
//go:generate rm -f ./duty_encoding.go
//go:generate rm -f ./messages_encoding.go
//go:generate rm -f ./share_encoding.go

//go:generate go run .../fastssz/sszgen --path messages.go --include ./message_id.go --exclude-objs MessageID,MsgType
//go:generate go run .../fastssz/sszgen --path duty.go
//go:generate go run .../fastssz/sszgen --path consensus_data.go --include ./duty.go,$GOPATH/pkg/mod/github.com/attestantio/go-eth2-client@v0.15.1/spec/altair,$GOPATH/pkg/mod/github.com/attestantio/go-eth2-client@v0.15.1/spec/bellatrix,$GOPATH/pkg/mod/github.com/attestantio/go-eth2-client@v0.15.1/spec/phase0,$GOPATH/pkg/mod/github.com/attestantio/go-eth2-client@v0.15.1/api/v1/bellatrix
//go:generate go run .../fastssz/sszgen --path share.go --include ./crypto.go --exclude-objs DomainType
