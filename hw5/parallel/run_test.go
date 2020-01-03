package parallel

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tasks := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task one")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task two")
			return errors.New("")
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task 3")
			return nil//errors.New("")
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task 4")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task 5")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("task 6")
			return nil
		},
	}

	Run(tasks, 2, 1)
}
