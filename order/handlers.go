package order

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SiparisVer(w http.ResponseWriter, r *http.Request) {
	// bir dışarıdan gelen veriyi tutacağımız nesneyi oluşturduk
	var requestBody struct {
		Description string `json:"description"`
		IsUser      bool   `json:"is_user"`
	}
	// dışarıdan gelen veriyi tuttuk
	json.NewDecoder(r.Body).Decode(&requestBody)

	if requestBody.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Lütfen sipariş açıklaması giriniz"))
		return
	}

	if !requestBody.IsUser {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sadece Bu hizmetten müşterilerimiz yararlanabilir. Lütfen öncelikli olarak kayıt olunuz."))
		return
	}
	// yeni bir sipariş oluşturduk
	siparis := NewSiparis(requestBody.Description)
	// yeni siparişin bilgilerini kullanıcı ile paylaştık
	json.NewEncoder(w).Encode(siparis)
}

func TumSiparisler(w http.ResponseWriter, r *http.Request) {
	siparisler := []Siparis{}

	for _, siparis := range Siparisler {
		fmt.Println(siparis.Code)
		siparisler = append(siparisler, *siparis)
	}
	json.NewEncoder(w).Encode(siparisler)
}

func Tamamlandi(w http.ResponseWriter, r *http.Request) {
	// sipariş id aldık
	siparisId := mux.Vars(r)["uuid"]

	// boş mu diye kontrol ettik
	if siparisId == "" {
		w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Lütfen sipariş kodu giriniz"))
        return
	}

	// sipariş bulunduysa tamalnadı olarak kaydettik. bulunamadısa bulunamadı diye döndürdük.
	if siparis, ok := Siparisler[siparisId]; ok {
		siparis.IsDelivered = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sipariş tamamlandı."))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sipariş kodu bulunamadı!"))
	}
}