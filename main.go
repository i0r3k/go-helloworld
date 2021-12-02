package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    // Disable Console Color
    // gin.DisableConsoleColor()
    gin.SetMode(gin.ReleaseMode)
    gin.DisableConsoleColor()

    r := gin.New()

    // Global middleware
    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    // By default gin.DefaultWriter = os.Stdout
    //r.Use(gin.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())

    // Ping test
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })

    return r
}

func main() {
    r := setupRouter()
    // Listen and Server in 0.0.0.0:18080
    r.Run(":18080")
}