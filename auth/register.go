package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	passowrd := c.PostForm("password")
	fmt.Println(name, email, passowrd)
}
