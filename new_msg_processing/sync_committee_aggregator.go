package new_msg_processing

func NewSyncCommitteeAggregatorRunner() *Runner {
	return &Runner{
		RoleType: "sync_committee_aggregator",
		config:   &RunnerConfg{},
	}
}
