package middlewares

import "github.com/gin-gonic/gin"

func UserAuth(c *gin.Context) {

	c.JSON(404, gin.H{"status": "200", "message": "Unauthorized user"})

}
