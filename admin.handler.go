package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os"
)

func showAdminPage(c *gin.Context) {
	err := qrcode.WriteFile(localAdr, qrcode.Medium, 256, "./assets/qr_code.png")
	if err != nil {
		os.Exit(1)
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Admin Page",
	})
}
