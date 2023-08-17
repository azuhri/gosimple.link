package controllers

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
	gee "github.com/tbxark/g4vercel"
)

func RootHandler(c *gee.Context) {
	c.JSON(http.StatusOK, H{
		"status":  true,
		"message": "success get data",
	})
}
