package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thepranavjain/bookstore_oauth-api/src/clients/cassandra"
	"github.com/thepranavjain/bookstore_oauth-api/src/domain/access_token"
	"github.com/thepranavjain/bookstore_oauth-api/src/http"
	"github.com/thepranavjain/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Making sure a cassandra session can be established
	cassandra.GetSession()
	fmt.Println("Cassandra DB Online.")

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}