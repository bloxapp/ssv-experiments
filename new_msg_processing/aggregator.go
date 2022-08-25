package new_msg_processing

type AggregatorRunner struct {
	RoleType string
	State    *State
}

func NewAggregatorRunner() Runner {
	return &AggregatorRunner{
		RoleType: "aggregator",
	}
}

func (r *AggregatorRunner) GetRole() string {
	return r.RoleType
}

func (r *AggregatorRunner) GetState() *State {
	return r.State
}

func (r *AggregatorRunner) StartNewDuty(duty interface{}) error {

}

func (r *AggregatorRunner) ProcessMsg(msg interface{}) error {

}
