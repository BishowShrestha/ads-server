package clicks

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, worker *ClickWorker) {
	r.POST("/ads/click", func(c *gin.Context) {
		var click ClickEvent
		if err := c.ShouldBindJSON(&click); err != nil {
			c.JSON(400, gin.H{"error": "invalid payload"})
			return
		}
		worker.Enqueue(click)
		c.JSON(200, gin.H{"status": "ok"})
	})
}
