package middleware

import (
	"github.com/EduardoNevesRamos/frelancer.git/common"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		const Bearer_schena = "Bearer "
		header := context.GetHeader("Authorization")
		if header == "" {
			context.AbortWithStatus(401)
		}

		token := header[len(Bearer_schena):]

		if !common.NewJWTService().ValidateToken(token) {
			context.AbortWithStatus(401)
		}
	}
}
