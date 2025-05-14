package main

import (
	"github.com/SisyphianLiger/Shootdev/api"
	"log"
	"net/http"
)

const filepathRoot = "./app"
const port = ":8080"

func main() {

	apiCfg := &api.ApiConfig{
		FileServerHits: 0,
		Jwt:            "Not Serviceable",
	}

	s := api.CreateNewServer()
	s.MountMiddleware(apiCfg)
	s.MakeApiHandlers(apiCfg)

	log.Fatal(http.ListenAndServe(port, s.Router))

}
