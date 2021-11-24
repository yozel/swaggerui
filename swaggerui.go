package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var dist embed.FS

var subFS fs.FS

func init() {
	var err error
	subFS, err = fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

}

func New(openApiSpec []byte) (http.Handler, error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(subFS)))
	mux.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write(openApiSpec) }))
	return mux, nil
}
