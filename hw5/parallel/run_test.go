package parallel

import (
	"errors"
	"sync/atomic"
	"testing"
)

func TestHandleTask(t *testing.T){
	input := make(chan func() error)
	output := make(chan error)
	go handleTask(input, output)

	input <- func() error {
		return nil
	}

	if err := <- output; err != nil{
		t.Errorf("Получено неожиоданное значение.")
	}

	input <- func() error {
		return errors.New("test")
	}

	if err := <- output; err == nil{
		t.Errorf("Получено неожиоданное значение.")
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

	if len(tasks) != int(taskCompleted){
		t.Errorf("Неверное количество выполненных задач.")
	}
}

func TestSendTask(t *testing.T){
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

	if !isCompleted{
		t.Errorf("Задача не была выполнена.")
	}
}