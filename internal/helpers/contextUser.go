package helpers

import "github.com/gin-gonic/gin"

const ContextUserKey = "auth_user_claims"

func GetUserFromContext(c *gin.Context) (*Claims, bool) {
	claims, exists := c.Get(ContextUserKey)
	if !exists {
		return nil, false
	}
	
	userClaims, ok := claims.(*Claims)
	if !ok {
		return nil, false
	}
	
	return userClaims, true
}

func MustGetUserFromContext(c *gin.Context) *Claims {
	user, ok := GetUserFromContext(c)
	if !ok {
		panic("user not found in context - auth middleware not applied?")
	}
	return user
}