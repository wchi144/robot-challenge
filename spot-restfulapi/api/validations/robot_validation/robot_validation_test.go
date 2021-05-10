package robot_validation

import (
	"spot-restfulapi/api/clients/robotclient"
	"spot-restfulapi/api/domain/robot_domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getStatusFunc func() robot_domain.RobotState
)

type robotClientMock struct{}

// Mock methods in robot client
func (mockClient *robotClientMock) EnqueueTask(commands string) (taskID string, position chan robot_domain.RobotState, err chan error) {
	return "", nil, nil
}

func (mockClient *robotClientMock) CancelTask(taskID string) error {
	return nil
}

func (mockClient *robotClientMock) CurrentState() robot_domain.RobotState {
	return getStatusFunc()
}

func testBase() {
	getStatusFunc = func() robot_domain.RobotState {
		return robot_domain.RobotState{
			X:        0,
			Y:        0,
			HasCrate: false,
		}
	}
	robotclient.RobotStruct = &robotClientMock{}
}

func IsCommandValid_PositiveXY_NoError(t *testing.T) {
	//Arrange
	testBase()

	//Act
	isValid, err := RobotValidator.IsCommandValid(strings.Repeat("N ", 10) + strings.Repeat(" E", 10))

	//Assert
	assert.True(t, isValid)
	assert.Nil(t, err)
}

func IsCommandValid_ExceedLimitXY_Error(t *testing.T) {
	//Arrange
	testBase()

	//Act
	isValid, err := RobotValidator.IsCommandValid(strings.Repeat("N ", 11) + strings.Repeat(" E", 11))

	//Assert
	assert.False(t, isValid)
	assert.NotNil(t, err)
}

func IsCommandValid_NegativeXY_Error(t *testing.T) {
	//Arrange
	testBase()

	//Act
	isValid, err := RobotValidator.IsCommandValid(strings.Repeat("W ", 1) + strings.Repeat(" S", 1))

	//Assert
	assert.False(t, isValid)
	assert.NotNil(t, err)
}
