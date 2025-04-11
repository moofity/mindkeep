package provider

import (
	"github.com/gin-gonic/gin"
)

type RoutesProvider struct {
	// Add any dependencies here (e.g., service, database connection)
}

func NewRoutesProvider() *RoutesProvider {
	return &RoutesProvider{}
}

func (ur *RoutesProvider) Provide(r *gin.Engine) {
	socialGraph := r.Group("/api/social-graph")
	{
		socialGraph.GET("/ping", ur.ping)
	}
}

// Handler methods
func (ur *RoutesProvider) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
