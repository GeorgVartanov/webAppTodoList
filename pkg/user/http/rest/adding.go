package rest

import (
	"github.com/GeorgVartanov/myWebApp/pkg/user/service/adding"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserHandler(s adding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newServiceUser, err := newUser.ConvertToAddingServiceUser()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = s.AddUser(*newServiceUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
		return
	}
}
