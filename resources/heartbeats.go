package resources

import (
	"net/http"
	"encoding/json"
	"github.com/italiviocorrea/golang/ibge/models"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.HeartbeatResponse{Status: "OK", Code: 200})

}

