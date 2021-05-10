package app

import "spot-restfulapi/api/controllers/robot_controller"

func routes() {
	router.POST("/robot", robot_controller.PostCommand)
	router.DELETE("/robot/:taskID", robot_controller.CancelCommand)
	router.GET("/robot/status", robot_controller.GetStatus)
}
