package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"tracka/pkg/config"

	"github.com/gorilla/mux"
)

// Culled from the mux docs
type spaHandler struct {
	staticPath, indexFile string
}

func (handler spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// unable to resolve absolute path
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend dirname to path
	path = filepath.Join(handler.staticPath, path)

	// check if dir contains files
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(handler.staticPath, handler.indexFile))
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(handler.staticPath)).ServeHTTP(w, r)
}

func Init(router *mux.Router) {

	api := router.PathPrefix("/api").Subrouter()

	spa := spaHandler{staticPath: config.Get().Static.StaticPath, indexFile: config.Get().Static.IndexPath}
	router.PathPrefix("/").Handler(spa)

	registerAuth(api)
}
