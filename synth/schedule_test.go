package synth

import (
	"testing"
	"time"
)

func Test_Schedule_Successful(t *testing.T) {
	c := make(chan Event)

	schedule := Schedule{
		Jobs: []JobSchedule{
			JobSchedule{
				Job:         MockJob{C: c},
				RepeatEvery: 3 * time.Second,
			},
		},
	}

	RunSchedule(schedule)

	select {
	case <-c:
		t.Log("OK")
	case <-time.After(5 * time.Second):
		t.Log("Schedule runner timed-out")
		t.Fail()
	}
}

func Test_Multiple_ScheduleJob(t *testing.T) {
	c := make(chan Event)

	schedule := Schedule{
		Jobs: []JobSchedule{
			JobSchedule{
				Job:         MockJob{C: c},
				RepeatEvery: 3 * time.Second,
			},
			JobSchedule{
				Job:         MockJob{C: c},
				RepeatEvery: 3 * time.Second,
			},
		},
	}

	RunSchedule(schedule)

	select {
	case <-c:
		t.Log("OK")
		select {
		case <-c:
			t.Log("OK")
		case <-time.After(8 * time.Second):
			t.Log("Schedule runner timed-out")
			t.Fail()
		}
	case <-time.After(8 * time.Second):
		t.Log("Schedule runner timed-out")
		t.Fail()
	}
}

type MockJob struct {
	C chan Event
}

func (m MockJob) Run() error {
	m.C <- Event{}
	return nil
}
