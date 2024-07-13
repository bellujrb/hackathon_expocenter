package middleware

import (
	"viva/db"
	_ "viva/docs"
	cashback "viva/internal/cashback/handler"
	mkt "viva/internal/mkt/handler"
	storages "viva/internal/storages/handler"
	user "viva/internal/user/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title
// @version 1.0
// @description API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(CORSConfig())
	r.Use(ResponseHandler())

	r.Use(func(c *gin.Context) {
		c.Set("db", db.Repo)
		c.Next()
	})
	//Use response, but not Token
	r.GET("/token", generateTokenHandler)

	auth := r.Group("/api")
	//auth.Use(authMiddleware)
	//Response and token service
	auth.PUT("/modify-cashback", storages.ModifyCashback)
	auth.PUT("/modify-pushed", storages.ModifyPushed)
	auth.PUT("/modify-used", storages.ModifyUse)
	auth.GET("/storages", storages.PullAllStorage)
	auth.GET("/storage/:id", storages.PullStorageId)

	auth.POST("/calc-cashback", cashback.CalculationCashBack)
	auth.POST("/remove-cashback", cashback.RemoveCashBack)

	auth.POST("/mkt", mkt.CreateItem)
	auth.GET("/all-item", mkt.AllItem)
	auth.GET("/mkt/:id", mkt.GetOneItem)

	auth.POST("/constumer", user.CreateConstumer)
	auth.GET("/all-constumer", user.AllConstumer)
	auth.GET("/constumer/:id", user.GetOneConstumer)
	auth.GET("/sync-customers", user.SyncConstumers)
	return r
}
