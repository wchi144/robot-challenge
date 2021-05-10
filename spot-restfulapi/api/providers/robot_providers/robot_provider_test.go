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

type getClientMock struct{}

// Mock methods in robot client
func (cm *getClientMock) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	return getRequestFunc(commands)
}

func (cm *getClientMock) CancelTask(taskID string) error {
	return cancelRequestFunc(taskID)
}

func (cm *getClientMock) CurrentState() robot_domain.RobotState {
	return getStatusFunc()
}

func TestEnqueueNoError(t *testing.T) {
	getRequestFunc = func(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
		return "123", make(chan robot_domain.RobotState), nil
	}

	// Mock robot client
	robotclient.RobotStruct = &getClientMock{}
	taskID, position, err := RobotProvider.EnqueueTask("N")
	assert.NotNil(t, taskID)
	assert.NotNil(t, position)
	assert.Nil(t, err)
}
