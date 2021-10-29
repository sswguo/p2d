package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  messages := getAllMessages()

  // Call the HTML method of the Context to render a template
  c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "index.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Home Page",
      "payload": messages,
    },
  )
}

func createMessage(c *gin.Context) {
  // handle form
  //title := c.PostForm("title")
  //nick := c.DefaultPostForm("nick", "anonymous")
  //content := c.PostForm("content")

  // handle JSON
  json := message{}

  c.BindJSON(&json)

  c.JSON(http.StatusOK, gin.H{
      "status":  gin.H{
          "status_code": http.StatusOK,
          "status":      "ok",
      },
      "title": json.Title,
      "content": json.Content,
  })
}