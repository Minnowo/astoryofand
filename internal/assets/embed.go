package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

//go:embed all:dist
var Assets embed.FS

func GetFileSystem(useOS bool) http.FileSystem {

	if useOS {

		log.Info("Using live assets")

		return http.FS(os.DirFS("app"))
	}

	log.Info("Using embed assets")

	fsys, err := fs.Sub(Assets, "dist")

	if err != nil {
		log.Fatal(err)
	}

	return http.FS(fsys)
}
