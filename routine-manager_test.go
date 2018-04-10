package routinero

import "testing"
import "time"

func TestDeadlockDetector(t *testing.T) {
	rm := new(RoutineManager)
	rm.Init(10000)
	time.Sleep(time.Second * 100)
}