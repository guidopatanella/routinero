package routinero

import "sync"
import "strings"
import "time"
import "runtime/pprof"
import "fmt"

type RoutineManager struct {
	mutex sync.Mutex
	RoutineMap map[int]Routine
}

func (rm *RoutineManager) Init(heartbeatMs int) {
	fmt.Println("routine manager - initialized")
	go rm.heartbeat(heartbeatMs)
}

func (rm *RoutineManager) heartbeat(heartbeat int) {
	fmt.Println("routine manager - heartbeat")
	rm.introspect()
	time.Sleep(time.Millisecond * time.Duration(heartbeat))
	go rm.heartbeat(heartbeat)
}

func (rm *RoutineManager) introspect() {
	pprof.Lookup("goroutine").WriteTo(rm, 3)
}

 func (rm *RoutineManager) Write(p []byte) (n int, err error) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	// (re)initialize routines map
	rm.RoutineMap = make(map[int]Routine)
	blockPos := 0
	routine := Routine{}
	text := string(p[:])
	row1 := "" // temp storage for intermediate row
	for i, row := range strings.Split(text, "\n") {
		if row=="" {
			blockPos = i+1
			// store routine
			fmt.Println("storing routine", routine.Number)
			rm.RoutineMap[routine.Number] = routine
			continue
		}
    	if i==blockPos {
			routine = Routine{}
			routine.ParseInit(row)
			fmt.Println("row", row)
		} else {
			if row1=="" {
				row1=row
			} else {
				routine.ParseStackRow(row1, row)
				row1=""
			}
		}
	}
	return len(p), nil
 }