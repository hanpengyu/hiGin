package route

import "github.com/gin-gonic/gin"

func RegisterRouters(engine *gin.Engine) {
	registerAdminRoute(engine)
}
