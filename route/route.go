package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is gcp " + action
	c.String(http.StatusOK, message)
}

func GetWelcome(c *gin.Context) {
	// /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest") // If not exists, return Guest in this case
	lastname := c.Query("lastname")                   // if not exists, return "" empty string
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
