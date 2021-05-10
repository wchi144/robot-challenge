//Mocking Robot SDK
package robotclient

import (
	"spot-restfulapi/api/domain/robot_domain"

	"github.com/google/uuid"
)

type robotStruct struct{}

type RobotInterface interface {
	EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() robot_domain.RobotState
}

var (
	RobotStruct RobotInterface = &robotStruct{}
)

func (r *robotStruct) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	taskGuid := uuid.New()
	var positionchan chan robot_domain.RobotState
	var errorchan chan error
	return taskGuid.String(), positionchan, errorchan
}

func (r *robotStruct) CancelTask(taskID string) error {
	return nil
}

func (r *robotStruct) CurrentState() robot_domain.RobotState {
	state := robot_domain.RobotState{
		X:        1,
		Y:        1,
		HasCrate: false,
	}
	return state
}
