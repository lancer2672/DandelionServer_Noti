package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lancer2672/DandelionServer_Noti/services"
	"github.com/lancer2672/DandelionServer_Noti/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ConfigServer(config utils.Config, s *services.NotificationService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("docs"))
	r.Handle("/docs/*", http.StripPrefix("/docs", fs))
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/docs/swagger.json"), //The url pointing to API definition
	))

	r.Get("/notification", func(w http.ResponseWriter, r *http.Request) {
		list, err := s.GetNotificationList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.JSONResponse(w, map[string]interface {
		}{
			"data":    list,
			"message": "Success",
		}, http.StatusOK)
	})
	return r
}
func StartServer(addr string, router http.Handler) {
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("Server start failed", err)
	}
	log.Println("Server started")
}
