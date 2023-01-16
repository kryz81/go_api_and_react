package handlers

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
)

func IndexHandler(frontendFs embed.FS) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := fs.ReadFile(frontendFs, "frontend/build/index.html")
		if err != nil {
			log.Fatal(err)
		}
		ctx.Header("Content-Type", "text/html")
		ctx.String(http.StatusOK, string(file))
	}
}
