package robot_controller

import (
	"net/http"
	"spot-restfulapi/api/domain/robot_domain"
	robot_service "spot-restfulapi/api/services/robot_service"

	"github.com/gin-gonic/gin"
)

func PostCommand(c *gin.Context) {
	var request robot_domain.RobotRequest
	c.BindJSON(&request)
	command := request.Commands
	taskId, _, _ := robot_service.RobotService.EnqueueTask(command)
	if len(taskId) <= 0 {
		c.JSON(500, gin.H{"error": "could not enqueue task"})
		return
	}
	c.JSON(http.StatusOK, taskId)
}

func CancelCommand(c *gin.Context) {
	taskID := c.Param("taskID")
	err := robot_service.RobotService.CancelTask(taskID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not cancel task"})
	}
	c.Status(http.StatusNoContent)
}

func GetStatus(c *gin.Context) {
	result := robot_service.RobotService.CurrentState()
	c.JSON(http.StatusOK, result)
}
