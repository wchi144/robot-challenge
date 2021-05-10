package robot_service

import (
	"errors"
	"spot-restfulapi/api/clients/robotclient"
	"spot-restfulapi/api/domain/robot_domain"
	"spot-restfulapi/api/validations/robot_validation"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc    func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error)
	cancelRequestFunc func(taskID string) error
	getStatusFunc     func() robot_domain.RobotState
	isValidCommandFun func(command string) (isValid bool, err error)
)

type robotClientMock struct{}
type robotValidatorMock struct{}

// Mock methods in robot client
func (mockClient *robotClientMock) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	return getRequestFunc(commands)
}

func (mockClient *robotClientMock) CancelTask(taskID string) error {
	return cancelRequestFunc(taskID)
}

func (mockClient *robotClientMock) CurrentState() robot_domain.RobotState {
	return getStatusFunc()
}

func (mockValidator *robotValidatorMock) IsCommandValid(command string) (isValid bool, err error) {
	return isValidCommandFun(command)
}

func TestEnqueueNoError(t *testing.T) {
	//Arrange
	getRequestFunc = func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
		return "123", make(chan robot_domain.RobotState), nil
	}
	isValidCommandFun = func(commands string) (isValid bool, err error) {
		return true, nil
	}

	robotclient.RobotStruct = &robotClientMock{}
	robot_validation.RobotValidator = &robotValidatorMock{}

	//Act
	taskID, position, err := RobotService.EnqueueTask("N")

	//Assert
	assert.NotNil(t, taskID)
	assert.NotNil(t, position)
	assert.Nil(t, err)
}

func TestEnqueueInvalidCommandReturnError(t *testing.T) {
	//Arrange
	getRequestFunc = func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
		return "123", make(chan robot_domain.RobotState), nil
	}
	isValidCommandFun = func(commands string) (isValid bool, err error) {
		return false, errors.New("Boo")
	}

	robotclient.RobotStruct = &robotClientMock{}
	robot_validation.RobotValidator = &robotValidatorMock{}

	//Act
	taskID, _, _ := RobotService.EnqueueTask(strings.Repeat("N ", 11))

	//Assert
	assert.Empty(t, taskID)
}
