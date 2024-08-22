package middleware

import (
	"net/http"
	"server/api/cookies"
	"server/api/handler"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(h handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the "auth_cookie"
		tokenString, err := c.Cookie("auth_cookie")
		if err != nil {
			// Handle the case where the authentication token is not present
			if c.Request.URL.Path == "/" {
				c.Redirect(http.StatusSeeOther, "/about")
				c.Abort() // Abort further processing
				return
			}
			c.Next()
			return
		}

		// Use decodeJWT to decode and validate the JWT
		userID, err := cookies.DecodeJWT(tokenString)
		if err != nil {
			// Handle the case where the authentication token is invalid
			if c.Request.URL.Path == "/" {
				c.Redirect(http.StatusSeeOther, "/about")
				c.Abort() // Abort further processing
				return
			}
			c.Next()
			return
		}

		// Fetch the user from the database using the userID
		user, err := h.Store.GetUserByID(c, userID)
		if err != nil {
			// If user retrieval fails, destroy the cookie and redirect to /login
			c.SetCookie("auth_cookie", "", -1, "/", "", false, true) // Invalidate the cookie
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // Abort further processing
			return
		}

		// Set the user in the context
		c.Set("user", user)

		// Proceed to the next handler
		c.Next()
	}
}
