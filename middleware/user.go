// package middleware

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // AuthMiddleware is a simple middleware example
// // that checks if a "user" exists in session (or header for simplicity)
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		// Example: check a header called "X-User" for simplicity
// 		user := c.GetHeader("X-User")

// 		if user == "" {
// 			// If no user found, block request
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized! Please login first.",
// 			})
// 			c.Abort() // Stop processing request
// 			return
// 		}

// 		// If user exists, continue to the next handler
// 		c.Next()
// 	}
// }
