package webRout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicRout struct{}

func (PublicRout) Load(engine *gin.Engine) {
	engine.StaticFS("/public", http.Dir("public"))
}
