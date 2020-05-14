package servers

import (
	"log"
	"net/http"
	"../constants"
	"strconv"
)

func StartServer(port int) {
	direccion := constants.TwoDots + strconv.Itoa(port)
	dir := http.Dir(constants.FolderDefault)
	pathFiles := http.FileServer(dir)
	path := http.StripPrefix(constants.PathAPI, pathFiles)

	
	http.Handle(constants.PathAPI, path)
	err := http.ListenAndServe(direccion, nil)
	if err != nil {
		log.Fatal(err)
	}
}
