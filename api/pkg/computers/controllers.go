package computers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matdorneles/gestao_ti/api/pkg/models"
)

// COMPUTER HANDLERS
// defines the computer handler
type ComputerHandler struct {
	computer computerAsset
}

// sets the interface and it's methods
type computerAsset interface {
	List() []models.Computer
	Get(id string) models.Computer
	Add(computer *models.Computer)
	Update(computer models.Computer)
	Delete(id string) models.Computer
}

// creates a new Computer Handler
func NewComputerHandler(ca computerAsset) *ComputerHandler {
	return &ComputerHandler{
		computer: ca,
	}
}

// gets all Computers
func (ch ComputerHandler) List(c *gin.Context) {
	computers := ch.computer.List()

	if len(computers) == 0 {
		c.JSON(http.StatusOK, gin.H{"result": "no records found"})
		return
	}

	c.JSON(http.StatusOK, computers)
}

// gets one Computer by ID
func (ch ComputerHandler) Get(c *gin.Context) {
	id := c.Param("id")

	computer := ch.computer.Get(id)

	if computer.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, computer)
}

// create a new Computer
func (ch ComputerHandler) Create(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch.computer.Add(&computer)
	c.JSON(http.StatusOK, computer)
}

// update an existing Computer
func (ch ComputerHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// find the Computer we want to Update
	computer := ch.computer.Get(id)

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
	ch.computer.Update(computer)
	c.JSON(http.StatusOK, computer)
}

// delete an existing Computer
func (ch ComputerHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	computer := ch.computer.Get(id)
	if computer.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	ch.computer.Delete(id)
	c.JSON(http.StatusOK, gin.H{"deleted": computer})
}
