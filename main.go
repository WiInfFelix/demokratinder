package main

import (
  "github.com/gin-gonic/gin"
  "github.com/pkg/browser"
  "net/http"
)

var router *gin.Engine
const tinderUrl = "https://www.tinder.com"
const localAdr = "http://localhost:8080/admin"

func main() {
  router = gin.Default()
  router.Static("/assets", "./assets")

  router.LoadHTMLGlob("templates/*")

  initializeRoutes()
  initBrowser()

  router.Run()
}

func render(c *gin.Context, data gin.H, templateName string){
  switch c.Request.Header.Get("Accept") {
  case "application/json":
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    c.XML(http.StatusOK, data["payload"])
  default:
    c.HTML(http.StatusOK, templateName, data)

  }
}

func initBrowser() {
  browser.OpenURL(tinderUrl)
  browser.OpenURL(localAdr)
}