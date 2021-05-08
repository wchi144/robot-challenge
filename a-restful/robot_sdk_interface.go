package main

import "fmt"

type Warehouse interface {
	Robots() []Robot
}

type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() RobotState
}

type RobotState struct {
	X        uint
	Y        uint
	HasCrate bool
}

func (r RobotState) GetString() string {
	return fmt.Sprintf("%d %d %t", r.X, r.Y, r.HasCrate)
}
