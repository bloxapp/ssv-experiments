package new_msg_processing

type Validator struct {
	Runners map[string]*Runner // runner, for each duty
}

func NewValidator() *Validator {
	return &Validator{
		Runners: map[string]Runner{
			"attester":                  NewAttesterRunner(),
			"proposer":                  NewProposerRunner(),
			"aggregator":                NewAggregatorRunner(),
			"sync_committee":            NewSyncCommitteeRunner(),
			"sync_committee_aggregator": NewSyncCommitteeAggregatorRunner(),
		},
	}
}

func (v *Validator) ProcessMessage(msg interface{}) error {
	runner := v.getRunnerForMsg(msg)
	return runner.ProcessMsg(msg)
}

func (v *Validator) getRunnerForMsg(msg interface{}) *Runner {
	panic("implement")
}
