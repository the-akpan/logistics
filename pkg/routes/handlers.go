package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"tracka/pkg/utils"
)

// Culled from the mux docs
type spaHandler struct {
	staticPath, indexFile string
}

func (handler spaHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path, err := filepath.Abs(req.URL.Path)
	if err != nil {
		// unable to resolve absolute path
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend dirname to path
	path = filepath.Join(handler.staticPath, path)

	// check if dir contains files
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(res, req, filepath.Join(handler.staticPath, handler.indexFile))
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(handler.staticPath)).ServeHTTP(res, req)
}

func notFound(res http.ResponseWriter, req *http.Request) {
	response := utils.CreateResponse(res)
	response.Message = "Not Found"
	res.WriteHeader(http.StatusNotFound)

	json.NewEncoder(res).Encode(response)
}
