package main

import (
	"fmt"
	provider "server/pkg/social-graph"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.DebugMode)

	provider.NewRoutesProvider().Provide(r)

	if err := r.Run(":8001"); err != nil {
		fmt.Println(err)
	}
}
