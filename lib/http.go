package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func RunHttpServer(port string) {
	http.HandleFunc("/promotions/", func(w http.ResponseWriter, r *http.Request) {
		getPromotion(w, r)
	})

	log.Printf("Http server is listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Printf("Can't start http server")
	}
}

func getPromotion(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/promotions/"):]

	promotion, err := GetPromotion(id)

	if err != nil {
		http.Error(w, "Can't get promotion", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(promotion)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func GetPromotion(id string) (Promotion, error) {
	_id, _ := strconv.Atoi(id)
	value, err := GetLineById(_id)

	if err != nil {
		return Promotion{}, err
	}

	return GetPromotionFromString(value), nil

}
