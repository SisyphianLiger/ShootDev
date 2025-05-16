package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SisyphianLiger/Shootdev/api"
	"github.com/SisyphianLiger/Shootdev/utils"
)

const filepathRoot = "./app"
const port = ":8080"

func main() {

	apiCfg := &api.ApiConfig{
		FileServerHits: 0,
		Jwt:            "Not Serviceable",
	}

	utils.OpenEnv()
	dbURL := utils.EnvironmentVarExists("DB_URL") // DB URL
	// devEnv := utils.EnvironmentVarExists("PLATFORM")      // Dev or Prod
	// jwtSecret := utils.EnvironmentVarExists("JWT_SECRET") // JWT

	// Make DB Connection extracted
	_ = utils.OpenDB("postgres", dbURL)
	fmt.Print("Made it here?\n")

	s := api.CreateNewServer()
	s.MountMiddleware(apiCfg)
	s.MakeApiHandlers(apiCfg)

	log.Fatal(http.ListenAndServe(port, s.Router))

}
