package robot_provider

import (
	"log"
	"spot-restfulapi/api/clients/robotclient"
	"spot-restfulapi/api/domain/robot_domain"
	"spot-restfulapi/api/validations/robot_validation"
)

type robotProvider struct{}

type robotServiceInterface interface {
	EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() robot_domain.RobotState
}

var (
	RobotProvider robotServiceInterface = &robotProvider{}
)

func (p *robotProvider) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	_, validationError := robot_validation.RobotValidator.IsCommandValid(commands)
	if validationError != nil {
		log.Fatal("Error: ", validationError)
		return "", nil, err
	}

	return robotclient.RobotStruct.EnqueueTask(commands)
}

func (p *robotProvider) CancelTask(taskID string) error {
	return robotclient.RobotStruct.CancelTask(taskID)
}

func (p *robotProvider) CurrentState() robot_domain.RobotState {
	return robotclient.RobotStruct.CurrentState()
}
