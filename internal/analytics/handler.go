package analytics

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, service *AnalyticsService) {
	r.GET("/ads/analytics", func(c *gin.Context) {
		data, err := service.GetClickCounts()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})
}
