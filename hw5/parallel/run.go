package parallel

func Run(task []func() error, N int, M int) {

	tasks := make(chan func() error)

	for i := 0; i < N; i++ {
		go func(chan<- func() error) {

		}(tasks)
	}
}
