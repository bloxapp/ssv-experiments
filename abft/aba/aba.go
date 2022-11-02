package aba

type ABA interface {
	// Run receives reliably received messages and returns true if decided with decided output
	Run(values map[int]bool) (bool, bool)
}
