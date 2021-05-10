package robot_validation

import (
	"errors"
	"fmt"
	"spot-restfulapi/api/clients/robotclient"
	"strings"
)

type robotValidator struct{}

type robotValidationInterface interface {
	IsCommandValid(commands string) (isValid bool, err error)
}

var (
	RobotValidator robotValidationInterface = &robotValidator{}
)

func (v *robotValidator) IsCommandValid(command string) (isValid bool, err error) {
	currentStatus := robotclient.RobotStruct.CurrentState()

	northCount := uint(strings.Count(command, "N"))
	southCount := uint(strings.Count(command, "S"))
	finalY := currentStatus.Y - southCount + northCount
	isYValid := 0 <= finalY && finalY <= 10
	if !isYValid {
		return false, errors.New(fmt.Sprintf("Invalid command. Cannot move %d unit on the y-axis", finalY))
	}

	eastCount := uint(strings.Count(command, "E"))
	westCount := uint(strings.Count(command, "W"))
	finalX := currentStatus.X - westCount + eastCount
	isXValid := 0 <= finalX && finalX <= 10
	if !isXValid {
		return false, errors.New(fmt.Sprintf("Invalid command. Cannot move %d unit on the x-axis", finalX))
	}

	return true, nil
}
