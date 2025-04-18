package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	bannerService := NewBannerService()

	http.HandleFunc("/user_banner", func(w http.ResponseWriter, r *http.Request) {
		userBannerHandler(w, r, bannerService)
	})
	http.HandleFunc("/banner", func(w http.ResponseWriter, r *http.Request) {
		bannerHandler(w, r, bannerService)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func bannerHandler(w http.ResponseWriter, r *http.Request, bannerService *BannerService) {
	if r.Method == http.MethodGet {
		adminToken := r.Header.Get("token")
		if adminToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		featureID, _ := strconv.Atoi(r.URL.Query().Get("feature_id"))
		tagID, _ := strconv.Atoi(r.URL.Query().Get("tag_id"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		// Получение всех баннеров с учетом фильтров
		banners, err := bannerService.getAllBanners(adminToken, featureID, tagID, limit, offset)
		if err != nil {
			http.Error(w, "Failed to get banners", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(banners)
		return
	}

	if r.Method == http.MethodPost {
		adminToken := r.Header.Get("token")
		if adminToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var bannerData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&bannerData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Создание нового баннера
		bannerID, err := bannerService.createBanner(adminToken, bannerData["tag_ids"].([]int), bannerData["feature_id"].(int), bannerData["content"].(map[string]interface{}), bannerData["is_active"].(bool))
		if err != nil {
			http.Error(w, "Failed to create banner", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"banner_id": bannerID})
		return
	}

	if r.Method == http.MethodPatch {
		adminToken := r.Header.Get("token")
		if adminToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		bannerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid banner ID", http.StatusBadRequest)
			return
		}

		var bannerData map[string]interface{}
		err = json.NewDecoder(r.Body).Decode(&bannerData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Обновление содержимого баннера
		err = bannerService.updateBanner(adminToken, bannerID, bannerData["tag_ids"].([]int), bannerData["feature_id"].(int), bannerData["content"].(map[string]interface{}), bannerData["is_active"].(bool))
		if err != nil {
			http.Error(w, "Failed to update banner", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodDelete {
		adminToken := r.Header.Get("token")
		if adminToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		bannerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid banner ID", http.StatusBadRequest)
			return
		}

		// Удаление баннера по идентификатору
		err = bannerService.deleteBanner(adminToken, bannerID)
		if err != nil {
			http.Error(w, "Failed to delete banner", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

