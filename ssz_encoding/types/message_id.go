package types

import (
	"encoding/binary"
)

const MessageIDSIze = 12

type MessageID [12]byte

var NoMessageID = [12]byte{}

func (msg MessageID) GetValidatorIndex() []byte {
	return msg[0:8]
}

func (msg MessageID) GetRoleType() BeaconRole {
	roleByts := msg[8:12]
	return BeaconRole(binary.LittleEndian.Uint32(roleByts))
}

func NewMsgID(validatorIndex uint64, role BeaconRole) MessageID {
	indexByts := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexByts, validatorIndex)

	roleByts := make([]byte, 4)
	binary.LittleEndian.PutUint32(roleByts, uint32(role))

	ret := MessageID{}
	copy(ret[:], append(indexByts, roleByts...))
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
