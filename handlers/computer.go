package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matdorneles/gestao_ti/config"
	"github.com/matdorneles/gestao_ti/models"
)

// COMPUTER HANDLERS

// get all Computers
func GetAllComputers(c *gin.Context) {
	var computers []models.Computer

	// Find finds all computers from database
	config.DB.Find(&computers)

	c.JSON(http.StatusOK, computers)
}

// get one Computer by ID
func GetComputer(c *gin.Context) {
	var computer models.Computer
	id := c.Param("id")

	// First finds the first result that matches the id from URL Param
	config.DB.First(&computer, id)

	// if computer.ID == 0 the result was not found
	if computer.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Computer not found"})
		return
	}

	c.JSON(http.StatusOK, computer)
}

// create a new Computer
func CreateNewComputer(c *gin.Context) {
	var computer models.Computer

	if err := c.BindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&computer)
	c.JSON(http.StatusCreated, computer)
}

// update an existing Computer
func PatchComputer(c *gin.Context) {
	var computer models.Computer
	id := c.Params.ByName("id")

	// find the Computer we want to Update
	config.DB.First(&computer, id)

	// if computer.ID == 0 the result was not found
	if computer.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Computer not found"})
		return
	}

	// save our new body into the found Computer
	if err := c.BindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the Computer with the new values
	config.DB.Model(&computer).UpdateColumns(computer)
	c.JSON(http.StatusOK, computer)
}

// delete an existing Computer
func DeleteComputer(c *gin.Context) {
	var computer models.Computer
	id := c.Param("id")
	config.DB.Delete(&computer, id)

	c.JSON(http.StatusOK, gin.H{"deleted": "successful"})
}
