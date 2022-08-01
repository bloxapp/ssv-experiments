package types

type MsgType [4]byte

var (
	// ConsensusProposeMsgType QBFT propose consensus message
	ConsensusProposeMsgType = MsgType{0x1, 0x0, 0x0, 0x0}
	// ConsensusPrepareMsgType QBFT prepare consensus message
	ConsensusPrepareMsgType = MsgType{0x1, 0x1, 0x0, 0x0}
	// ConsensusCommitMsgType QBFT commit consensus message
	ConsensusCommitMsgType = MsgType{0x1, 0x2, 0x0, 0x0}
	// ConsensusRoundChangeMsgType QBFT round change consensus message
	ConsensusRoundChangeMsgType = MsgType{0x1, 0x3, 0x0, 0x0}

	// DecidedMsgType are all QBFT decided messages
	DecidedMsgType = MsgType{0x2, 0x0, 0x0, 0x0}

	// PartialSignatureMsgType are all partial signatures msgs over beacon chain specific signatures
	PartialSignatureMsgType = MsgType{0x3, 0x0, 0x0, 0x0}

	// DKGInitMsgType sent when DKG instance is started by requester
	DKGInitMsgType = MsgType{0x4, 0x0, 0x0, 0x0}
	// DKGProtocolMsgType contains all key generation protocol msgs
	DKGProtocolMsgType = MsgType{0x4, 0x1, 0x0, 0x0}
	// DKGDepositDataMsgType post DKG deposit data signatures
	DKGDepositDataMsgType = MsgType{0x4, 0x2, 0x0, 0x0}
	// DKGOutputMsgType final output msg used by requester to make deposits and register validator with SSV
	DKGOutputMsgType = MsgType{0x4, 0x3, 0x0, 0x0}

	// UnknownMsgType can't be identified
	UnknownMsgType = MsgType{0x0, 0x0, 0x0, 0x0}
)

type Message struct {
	ID            MessageID `ssz-size:"32"`
	DataSSZSnappy []byte    `ssz-max:"2048"`
}
