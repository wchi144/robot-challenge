package robot_service

import (
	"fmt"
	"spot-restfulapi/api/clients/robotclient"
	"spot-restfulapi/api/domain/robot_domain"
	"spot-restfulapi/api/validations/robot_validation"
)

type robotService struct{}

type robotServiceInterface interface {
	EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() robot_domain.RobotState
}

var (
	RobotService robotServiceInterface = &robotService{}
)

func (p *robotService) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	_, validationError := robot_validation.RobotValidator.IsCommandValid(commands)
	if validationError != nil {
		fmt.Println("Failed to enqueueTask. Error: " + validationError.Error())
		return "", nil, err
	}

	return robotclient.RobotStruct.EnqueueTask(commands)
}

func (p *robotService) CancelTask(taskID string) error {
	return robotclient.RobotStruct.CancelTask(taskID)
}

func (p *robotService) CurrentState() robot_domain.RobotState {
	return robotclient.RobotStruct.CurrentState()
}
