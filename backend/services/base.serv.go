package services

import (
	"new-fix/models"

	"github.com/gin-gonic/gin"
)

type (
	BaseService struct {
		Model      *models.MySqlMdl
		Ctx        *gin.Context
		DbConnName string
	}
)
