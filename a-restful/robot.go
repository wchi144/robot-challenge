package main

type Spot struct {
}

func (s Spot) EnqueueTask(commands string) (string, chan, chan) {
	taskID := "123"
	var positionchan chan RobotState
	var errorchan chan error
	return taskID, positionchan, errorchan
}

func (s Spot) CancelTask(taskID string) chan {
	var errorchan chan error
	return errorchan
}

func (s Spot) CurrentState() RobotState {
	state := RobotState{
		X:        1,
		Y:        1,
		HasCrate: true,
	}
	return state
}


