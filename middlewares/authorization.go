package middlewares

import (
	"sosialMedia/helpers"

	"github.com/gin-gonic/gin"
)

func CanAccess(user_id int, c *gin.Context) bool {
	flag := false
	if user_id == helpers.ClaimToken(c).ID || helpers.ClaimToken(c).Role == "Admin" {
		flag = true
	}

	return flag
}
