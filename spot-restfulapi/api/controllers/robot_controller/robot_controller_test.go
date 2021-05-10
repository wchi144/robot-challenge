package robot_controller

import (
	"net/http"
	"net/http/httptest"
	"spot-restfulapi/api/domain/robot_domain"
	robot_provider "spot-restfulapi/api/providers/robot_providers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	createTaskFunc func(command string) (taskID string, position chan robot_domain.RobotState, err chan error)
	cancelTaskFunc func(taskID string) error
	getStatusFunc  func() robot_domain.RobotState
)

type robotProviderMock struct{}

func (r *robotProviderMock) EnqueueTask(command string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	return createTaskFunc(command)
}

func (r *robotProviderMock) CancelTask(taskID string) error {
	return cancelTaskFunc(taskID)
}

func (r *robotProviderMock) CurrentState() robot_domain.RobotState {
	return getStatusFunc()
}

func TestGetStatus(t *testing.T) {
	//Arrange
	getStatusFunc = func() robot_domain.RobotState {
		return robot_domain.RobotState{
			X:        4,
			Y:        5,
			HasCrate: false,
		}
	}
	robot_provider.RobotProvider = &robotProviderMock{}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)

	//Act
	GetStatus(c)

	//Assert
	assert.EqualValues(t, http.StatusOK, response.Code)
}
