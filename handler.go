package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func handleTake(c *gin.Context) {
	var json GameRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var robotTake = Take(rand.Intn(3))
	var userResult = json.UserTake.Match(robotTake)

	c.JSON(http.StatusOK, GameResponse{
		RobotTake:   robotTake,
		RobotResult: robotTake.Match(json.UserTake),
		UserTake:    json.UserTake,
		UserResult:  userResult,
	})
}
