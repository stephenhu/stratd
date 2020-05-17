package main

import (
	"encoding/json"
  "net/http"

)

func versionHandler(w http.ResponseWriter, r *http.Request) {

  switch r.Method {
  case http.MethodGet:

		v := map[string]string{
			"version": APP_VERSION,
		}

		j, jsonErr := json.Marshal(v)

		if jsonErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(j)
		}


  case http.MethodPost:
  case http.MethodDelete:
	case http.MethodPut:
	default:
	  w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // versionHandler
