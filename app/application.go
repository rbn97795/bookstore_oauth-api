package app

import (
	"bookstore_oauth-api/domain/access_token"
	"bookstore_oauth-api/http"
	"bookstore_oauth-api/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHander := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHander.GetById)

	router.Run(":8080")
}
