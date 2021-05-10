package robot_controller

import (
	"net/http"
	robot_provider "spot-restfulapi/api/providers/robot_providers"

	"github.com/gin-gonic/gin"
)

func PostCommand(c *gin.Context) {
	command := c.PostForm("command")
	taskId, _, err := robot_provider.RobotProvider.EnqueueTask(command)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not enqueue task"})
	}
	c.JSON(http.StatusOK, taskId)
}

func CancelCommand(c *gin.Context) {
	taskID := c.Param("taskID")
	err := robot_provider.RobotProvider.CancelTask(taskID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not cancel task"})
	}
	c.Status(http.StatusNoContent)
}

func GetStatus(c *gin.Context) {
	result := robot_provider.RobotProvider.CurrentState()
	c.JSON(http.StatusOK, result)
}
