package server

import (
	"log"
	"net/http"

	"github.com/lancer2672/DandelionServer_Noti/services"
	"github.com/lancer2672/DandelionServer_Noti/utils"
)

func ConfigServer(config utils.Config, s *services.NotificationService) {
	http.HandleFunc("/notification", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			list, err := s.GetNotificationList()
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			utils.JSONResponse(res, map[string]interface {
			}{
				"data":    list,
				"message": "Success",
			}, http.StatusOK)
		} else {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/", func(r http.ResponseWriter, req *http.Request) {
		log.Println("no thing")
	})
}
func StartServer(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Server start failed", err)
	}
	log.Println("Server started")
}
