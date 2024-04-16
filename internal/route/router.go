package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	gin.SetMode(gin.DebugMode)
	ginEngine := gin.New()

	ginEngine.Use(corsMiddleware())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(gin.Logger())

	api := ginEngine.Group("/api")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
			ctx.String(http.StatusOK, "ok")
		})

		if mode := gin.Mode(); mode == gin.DebugMode {
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}