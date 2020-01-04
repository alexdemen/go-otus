package parallel

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestHandleTask(t *testing.T) {
	input := make(chan func() error)
	output := make(chan error)
	go handleTask(input, output)

	input <- func() error {
		return nil
	}

	if err := <-output; err != nil {
		t.Errorf("Получено неожиоданное значение.")
	}

	input <- func() error {
		return errors.New("test")
	}

	if err := <-output; err == nil {
		t.Errorf("Получено неожиоданное значение.")
	}
}

func TestSendTask(t *testing.T) {
	var temp int
	var isCompleted bool
	taskChan := make(chan func() error)
	task := func() error {
		isCompleted = true
		return nil
	}
	go sendTask(&temp, task, taskChan)

	receivedTask := <-taskChan
	receivedTask()

	if !isCompleted {
		t.Errorf("Задача не была выполнена.")
	}
}

func TestRun(t *testing.T) {
	var taskCompleted uint32
	task := func() error {
		atomic.AddUint32(&taskCompleted, 1)
		return nil
	}

	tasks := []func() error{task, task, task, task, task, task,
	}

	Run(tasks, 2, 1)

	if len(tasks) != int(taskCompleted) {
		t.Errorf("Неверное количество выполненных задач.")
	}
}

func TestRunWithError(t *testing.T) {
	var taskCompleted uint32
	var errorCount uint32
	task := func() error {
		atomic.AddUint32(&taskCompleted, 1)
		return nil
	}
	errTask := func() error {
		atomic.AddUint32(&taskCompleted, 1)
		atomic.AddUint32(&errorCount, 1)
		time.Sleep(1 * time.Millisecond)
		return errors.New("")
	}

	tasks := []func() error{errTask, errTask, task, task, task, task, task}

	N := 2
	M := 1

	Run(tasks, N, M)

	if int(errorCount) != 2 {
		t.Errorf("Неверное количество ошибок в выполненных задачах.")
	}
	if int(taskCompleted) > N+M {
		t.Errorf("Количество выполненых задач превышает N+M=%d", N+M)
	}

}
