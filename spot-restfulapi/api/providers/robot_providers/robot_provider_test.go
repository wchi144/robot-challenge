package robot_provider

import (
	"spot-restfulapi/api/clients/robotclient"
	"spot-restfulapi/api/domain/robot_domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc    func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error)
	cancelRequestFunc func(taskID string) error
	getStatusFunc     func() robot_domain.RobotState
)

type robotClientMock struct{}

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

func TestEnqueueNoError(t *testing.T) {
	//Arrange
	getRequestFunc = func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
		return "123", make(chan robot_domain.RobotState), nil
	}
	robotclient.RobotStruct = &robotClientMock{}

	//Act
	taskID, position, err := RobotProvider.EnqueueTask("N")

	//Assert
	assert.NotNil(t, taskID)
	assert.NotNil(t, position)
	assert.Nil(t, err)
}
