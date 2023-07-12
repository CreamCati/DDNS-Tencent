package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
