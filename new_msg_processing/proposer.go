package new_msg_processing

func NewProposerRunner() *Runner {
	return &Runner{
		RoleType: "proposer",
		config:   &RunnerConfg{},
	}
}
