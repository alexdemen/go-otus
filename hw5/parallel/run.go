package parallel

func Run(task []func() error, N int, M int) {
	runTasksProcessing(task, N, M)
}

func runTasksProcessing(tasks []func() error, parallelTasksCount, maxErrCount int) {

	taskChan := make(chan func() error)
	taskCompleted := make(chan error)

	for i := 0; i < parallelTasksCount; i++ {
		go handleTask(taskChan, taskCompleted)
	}

	processTasks(taskChan, taskCompleted, maxErrCount, tasks)
}

func processTasks(taskChan chan func() error, taskCompleted chan error,
	maxErrCount int, tasks []func() error) {

	var errCount, taskCount int
	var curTaskId int
	var isEnd bool

taskLoop:
	for {
		select {
		case err := <-taskCompleted:
			if err != nil {
				errCount++
			}
			taskCount++

			if !isEnd && (errCount >= maxErrCount || taskCount == len(tasks)) {
				close(taskChan)
				isEnd = true
			}

			if taskCount == curTaskId && isEnd {
				break taskLoop
			}

		default:
			if curTaskId < len(tasks) && !isEnd {
				sendTask(&curTaskId, tasks[curTaskId], taskChan)
			}
		}
	}

}

func sendTask(curTaskId *int, task func() error, taskChan chan func() error) {
	select {
	case taskChan <- task:
		*curTaskId++
	default:
	}
}

func handleTask(input <-chan func() error, output chan<- error) {
	for {
		if task, ok := <-input; ok {
			output <- task()
		} else {
			break
		}
	}
}
