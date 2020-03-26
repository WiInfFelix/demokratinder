package main

func initializeRoutes() {
  router.GET("/", showIndexPage)
  router.GET("/article/view/:article_id", getArticle)

  router.GET("/admin", showAdminPage)
  router.GET("/voting", votingPage)

  router.POST("/voting", votingPage)
}
