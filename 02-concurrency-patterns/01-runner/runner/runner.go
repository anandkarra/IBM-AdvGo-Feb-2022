/* implement the runner */
package runner

import (
	"fmt"
	"sync"
	"time"
)

type RunnerImpl interface {
	New(timeout time.Duration) Runner
	Add(func(int))
	Start()
}

type Task func(int)

type Runner struct {
	Timeout   time.Duration
	Tasks     []Task
	Wg        sync.WaitGroup
	StartTime time.Time
}

var buffer = time.Duration(1 * time.Second)

// var wg sync.WaitGroup

func New(timeout time.Duration) *Runner {
	fmt.Printf("Initializing runnner, timeout=%v\n", timeout)

	var r Runner
	r.Timeout = timeout
	r.Tasks = []Task{}
	// r.Wg = wg
	r.StartTime = time.Now()

	fmt.Println("Runner initialized")
	return &r
}

func (r *Runner) Add(fn Task) {
	fmt.Printf("Adding task... ")

	r.Wg.Add(1)
	r.Tasks = append(r.Tasks, fn)
	// r.Tasks[id] = fn(id)

	fmt.Printf("✅ Task added\n")
}

func (r *Runner) Start() {
	fmt.Printf("Total # of tasks=%d\n", len(r.Tasks))

	for taskNum, task := range r.Tasks {
		fmt.Printf("Running task #%d\n", taskNum+1)

		go task(taskNum)
		wg.Done()
		if time.Since(r.StartTime) > r.Timeout {
			fmt.Println("❌ Timeout")
		}
	}
	if time.Since(r.StartTime) < r.Timeout {
		fmt.Println("✅ Success")
	}
}
