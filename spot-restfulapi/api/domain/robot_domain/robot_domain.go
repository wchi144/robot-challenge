package robot_domain

type RobotState struct {
	X        uint `json:"x"`
	Y        uint `json:"y"`
	HasCrate bool `json:"hasCrate"`
}

type RobotRequest struct {
	Commands string `json:"command"`
}
