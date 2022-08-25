package new_msg_processing

type RunnerConfg struct {
	PreConsensusF         func(signedMsg interface{}) error
	ValidatePreConsensusF func(signedMsg interface{}) error
	SignPostConsensusF    func(decidedValue, signer interface{}) error
	executeDutyF          func(duty interface{}) error
}

type Runner interface {
	GetRole() string
	GetState() *State
	StartNewDuty(duty interface{}) error
	ProcessMsg(msg interface{}) error
}

//type runner struct {
//
//}
//
//func (dr *runner) StartNewDuty(duty interface{}) error {
//	return dr.executeDuty(duty)
//}
//
//func (dr *runner) ProcessMsg(msg interface{}) error {
//	if err := dr.validateMsg(msg); err != nil {
//		return err
//	}
//
//	switch msg.(string) {
//	case "consensus":
//		return dr.processConsensusMessage(msg)
//	case "partial_sig":
//		isPostConsensusMsg := true // should be a real check
//		if isPostConsensusMsg {
//			return dr.processPostConsensusMessage(msg)
//		}
//		return dr.processPreConsensusMessage(msg)
//	default:
//		return fmt.Errorf("unknown msg")
//	}
//}
//
//func (dr *runner) canStartNewDuty()                                  {}
//func (dr *runner) signBeaconObject()                                 {}
//func (dr *runner) validateMsg(msg interface{}) error                 { panic("implement") }
//func (dr *runner) decide()                                           {}
//func (dr *runner) processConsensusMessage(msg interface{}) error     { panic("implement") }
//func (dr *runner) processPostConsensusMessage(msg interface{}) error { panic("implement") }
//
//func (dr *runner) processPreConsensusMessage(signedMsg interface{}) error {
//	return dr.config.PreConsensusF(signedMsg)
//}
//func (dr *Runner) validatePreConsensusMsg(signedMsg interface{}) error {
//	return dr.config.ValidatePreConsensusF(signedMsg)
//}
//func (dr *Runner) signDutyPostConsensus(decidedValue, signer interface{}) error {
//	return dr.config.SignPostConsensusF(decidedValue, signer)
//}
//func (dr *Runner) executeDuty(duty interface{}) error {
//	return dr.config.executeDutyF(duty)
//}
