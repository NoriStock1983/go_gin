package commons

import (
	"go_gin/main/api/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/test", controllers.TestController)

	r.Run()
}
