package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler(c *gin.Context) {
	htmlStr := `<html>
	<head></head>
	<body>
	<center>
	<h1>Hey GopherCon UK!</h1>
	<p><img src="/img/gopher-golfer.png"/></p>
	<p>What better way to demo Athens than with cat pictures?</p>
	<h1>A demo with cat <i>and</i> dog pictures</h1>
	<p><a href="/kitty">Cats</a></p>
	<p><a href="/pup">Dogs</a></p>
	</center>
	</body>
	</html>`
	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}
