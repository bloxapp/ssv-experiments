package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
)

const MessageIDSIze = 32

type MessageID [32]byte

var NoMessageID = [32]byte{}

func (msg MessageID) GetValidatorIndex() uint64 {
	roleByts := msg[0:8]
	return binary.LittleEndian.Uint64(roleByts)
}

func (msg MessageID) GetRoleType() BeaconRole {
	roleByts := msg[8:12]
	return BeaconRole(binary.LittleEndian.Uint32(roleByts))
}

func (msg MessageID) GetETHAddress() common.Address {
	ret := common.Address{}
	copy(ret[:], msg[0:20])
	return ret
}

func (msg MessageID) GetDKGIndex() uint32 {
	roleByts := msg[20:24]
	return binary.LittleEndian.Uint32(roleByts)
}

func (msg MessageID) GetMsgType() MsgType {
	ret := MsgType{}
	copy(ret[:], msg[28:32])
	return ret
}

func NewMsgIDValidator(validatorIndex uint64, role BeaconRole, msgType MsgType) MessageID {
	indexByts := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexByts, validatorIndex)

	roleByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(roleByts, uint32(role))

	ret := MessageID{}
	copy(ret[:12], append(indexByts, roleByts...))
	copy(ret[28:], msgType[:])
	return ret
}

func NewMsgIDETHAddress(address common.Address, index uint32, msgType MsgType) MessageID {
	roleByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(roleByts, index)

	ret := MessageID{}
	copy(ret[:24], append(address[:], roleByts...))
	copy(ret[28:], msgType[:])
	return ret
}

type MessageBytes []byte

func (msgByts MessageBytes) MsgID() MessageID {
	if len(msgByts) < MessageIDSIze {
		return NoMessageID
	}

	ret := MessageID{}
	copy(ret[:], msgByts[0:MessageIDSIze])
	return ret
}

func (msgByts MessageBytes) MsgType() MsgType {
	if len(msgByts) < (MessageIDSIze + 4) {
		return UnknownMsgType
	}

	ret := MsgType{}
	copy(ret[:], msgByts[MessageIDSIze:MessageIDSIze+4])
	return ret
}
