package synth

// Job describes a synthetic transaction job to be performed
type Job interface {
	Run() error
}
