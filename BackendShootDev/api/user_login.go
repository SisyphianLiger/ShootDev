package api

import (
	"encoding/json"
	"github.com/SisyphianLiger/Shootdev/utils"
	"net/http"
	"time"
)

func (cfg *ApiConfig) userLogin(w http.ResponseWriter, r *http.Request) {
	// Process Email Login
	type parameters struct {
		Email            string
		Password         string
		ExpiresInSeconds *time.Duration `json:"expires_in_seconds,omitempty"`
	}

	p := parameters{}
	decoder := json.NewDecoder(r.Body)
	parameterErr := decoder.Decode(&p)
	if parameterErr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Body Not Correct", parameterErr)
		return
	}

}
