package robot_controller

import (
	"fmt"
	"net/http"
	"spot-restfulapi/api/domain/robot_domain"
	robot_provider "spot-restfulapi/api/providers/robot_providers"

	"github.com/gin-gonic/gin"
)

func PostCommand(c *gin.Context) {
	var request robot_domain.RobotRequest
	c.BindJSON(&request)
	command := request.Commands
	taskId, _, errorCh := robot_provider.RobotProvider.EnqueueTask(command)
	select {
	case err := <-errorCh:
		fmt.Println("Error: ", err)
		c.JSON(500, gin.H{"error": "could not enqueue task. Error: "})
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
