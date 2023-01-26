package handlers

import (
	"dream/domain/models"
	"dream/usecases/interactor"

	"github.com/gin-gonic/gin"
)

type dreamHandler struct {
	dreamInteractor interactor.DreamInteractor
}

type DreamHandler interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
}

func NewDreamHandler(i interactor.DreamInteractor) DreamHandler {
	return &dreamHandler{i}
}

func Greet(c *gin.Context) {
	c.JSON(200, "Hi Dreamer")
}

func (h *dreamHandler) Add(c *gin.Context) {
	var dm models.DreamModel
	c.ShouldBindJSON(&dm)
	if err := h.dreamInteractor.Add(c, &dm); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"result": "error"})
		return
	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}

func (h *dreamHandler) GetAll(c *gin.Context) {
	result, err := h.dreamInteractor.GetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"result": "error"})
		return
	}
	c.JSON(200, gin.H{
		"result": "succes",
		"object": result,
	})
}
