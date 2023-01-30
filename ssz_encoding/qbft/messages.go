package qbft

// Message includes the full consensus input to be decided on, used for decided, proposal and round-change messages
type Message struct {
	MsgType    uint32
	Height     uint64
	Round      uint64
	Identifier [52]byte `ssz-size:"52"`
	Data       []byte   `ssz-max:"2048"`
}

// SignedMessage includes a signature over Message AND optional justification fields (not signed over)
type SignedMessage struct {
	Message   Message
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

//func (msg *SignedMessage) Validate() error {
//	if msg.Object != nil {
//		r, err := msg.Object.HashTreeRoot()
//		if err != nil {
//			return errors.Wrap(err, "could not get object root")
//		}
//		if !bytes.Equal(msg.Message.Root[:], r[:]) {
//			return errors.Wrap(err, "object root not equal to message root")
//		}
//	}
//	return nil
//}
//
//type Justifications struct {
//	RoundChangeJustifications [][]byte `ssz-max:"13,1024"`
//	ProposalJustifications    [][]byte `ssz-max:"13,1024"`
//}
//
//func (j *Justifications) GetRoundChangeJustifications() ([]*SignedMessage, error) {
//	return j.toSignedMessages(j.RoundChangeJustifications)
//}
//
//func (j *Justifications) GetProposalJustifications() ([]*SignedMessage, error) {
//	return j.toSignedMessages(j.ProposalJustifications)
//}
//
//func (j *Justifications) toSignedMessages(data [][]byte) ([]*SignedMessage, error) {
//	ret := make([]*SignedMessage, len(data))
//	for i, byts := range data {
//		msg := &SignedMessage{}
//		if err := msg.UnmarshalSSZ(byts); err != nil {
//			return nil, errors.Wrap(err, "could not unmarshal signed message")
//		}
//		ret[i] = msg
//	}
//	return ret, nil
//}
