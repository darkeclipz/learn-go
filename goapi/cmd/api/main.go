package api

import (
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
