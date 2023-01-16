package utils

import (
	"embed"
	"io/fs"
	"log"
)

func GetFrontendAssetsRoute(frontendFs embed.FS) fs.FS {
	files, err := fs.Sub(frontendFs, "frontend/build/static")
	if err != nil {
		log.Fatal(err)
	}
	return files
}
