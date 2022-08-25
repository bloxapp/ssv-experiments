package new_msg_processing

func NewSyncCommitteeRunner() *Runner {
	return &Runner{
		RoleType: "sync_committee",
		config:   &RunnerConfg{},
	}
}
