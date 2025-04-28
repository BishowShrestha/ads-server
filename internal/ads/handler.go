package ads

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, service *AdService) {
	r.GET("/ads", func(c *gin.Context) {
		ads, err := service.ListAds()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, ads)
	})
}
