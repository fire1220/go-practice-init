package route

import (
	"github.com/gin-gonic/gin"
	"go-project/app/route/backend"
	"go-project/app/route/frontend"
)

func Register(r *gin.Engine) {
	backend.Backend(r)
	frontend.Register(r)
}
