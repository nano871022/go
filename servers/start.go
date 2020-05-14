package servers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../constants"
	"../operations"
)

func StartServer(port int) {
	direccion := constants.TwoDots + strconv.Itoa(port)
	dir := http.Dir(constants.FolderDefault)
	pathFiles := http.FileServer(dir)
	path := http.StripPrefix(constants.PathAPI, pathFiles)

	http.HandleFunc(constants.PathAPI+"suma", servicioSumar)
	http.Handle(constants.PathAPI, path)
	err := http.ListenAndServe(direccion, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func servicioSumar(response http.ResponseWriter, request *http.Request) {
	listaValores, existe := request.URL.Query()["valor"]
	if existe {
		x, _ := strconv.ParseInt(listaValores[0], 10, 8)
		y, _ := strconv.ParseInt(listaValores[1], 10, 8)
		fmt.Print(x)
		fmt.Print(y)
		valor := operations.Suma(int8(x), int8(y))
		fmt.Fprintf(response, "Valor sumado :: %d", valor)
	}

}
