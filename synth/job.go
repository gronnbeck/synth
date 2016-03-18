package synth

import "time"

// Job describes a synthetic transaction job to be performed
type Job interface {
	Schedule() time.Duration
	Run() error
}
