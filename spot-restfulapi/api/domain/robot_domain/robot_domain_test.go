package robot_domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobotState(t *testing.T) {
	request := RobotState{
		X:        1,
		Y:        2,
		HasCrate: true,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var result RobotState
	err = json.Unmarshal(bytes, &result)

	assert.Nil(t, err)
	assert.EqualValues(t, result.X, request.X)
	assert.EqualValues(t, result.Y, request.Y)
	assert.EqualValues(t, result.HasCrate, request.HasCrate)
}
