package new_msg_processing

type AttesterRunner struct {
	RoleType string
	State    *State
}

func NewAttesterRunner() Runner {
	return &AttesterRunner{
		RoleType: "attester",
	}
}

func (r *AttesterRunner) GetRole() string {
	return r.RoleType
}

func (r *AttesterRunner) GetState() *State {
	return r.State
}

func (r *AttesterRunner) StartNewDuty(duty interface{}) error {

}

func (r *AttesterRunner) ProcessMsg(msg interface{}) error {

}
