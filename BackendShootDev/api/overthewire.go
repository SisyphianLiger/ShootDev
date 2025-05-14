package api

import (
	"encoding/json"
	"github.com/SisyphianLiger/Shootdev/utils"
	"net/http"
)

type Data struct {
	Body string `json:"data"`
}

func (api *ApiConfig) OverTheWire(w http.ResponseWriter, r *http.Request) {

	data := Data{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		// Make a utils folder that then we can import for JSON parsing
		utils.RespondWithError(w, 400, "The Payload is incorrectly specified", err)
		return
	}

	if data.Body == "" {
		utils.RespondWithError(w, http.StatusNoContent, "No body in the Payload", err)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, data.Body)
}
