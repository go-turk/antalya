package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-turk/antalya/order"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//ilksiparisdenemeleri()
	// Birinci : Bir Router Tanımlayalım
	r := mux.NewRouter()
	r.HandleFunc("/siparis/ver", SiparisVer).Methods("POST")
	r.HandleFunc("/siparisler", TumSiparisler).Methods("GET")
	fmt.Println(":9096 çalışmaya başladı")
	http.ListenAndServe(":9096", r)
	// - Sipariş ekleme
	// - Sipariş teslim etme
	// 1. adımda inMemory (Projenin içinde bir değişkene kaydedeceğiz)
	// 2. adımda bunları postgreSQL denilen Database'e kaydedeceğiz.
}

func ilksiparisdenemeleri() {
	ilkSiparis := order.NewSiparis("2 ekmek 1 çay")
	fmt.Println("siparis id:", ilkSiparis.Code)
	fmt.Println("siparis açıklaması:", ilkSiparis.Description)
	if ilkSiparis.IsDelivered {
		fmt.Println("Siraiş Teslim edildi")
	} else {
		fmt.Println("Sipariş henüz Teslim edilmedi")
	}
	order.Siparisler[ilkSiparis] = true
	fmt.Println(order.Siparisler)
}

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
	siparis := order.NewSiparis(requestBody.Description)
	// yeni siparişin bilgilerini kullanıcı ile paylaştık
	json.NewEncoder(w).Encode(siparis)
}

func TumSiparisler(w http.ResponseWriter, r *http.Request) {
	siparisler := []order.Siparis{}
	for siparis, durum := range order.Siparisler {
		fmt.Println(siparis.Code)
		if durum {
			siparisler = append(siparisler, *siparis)
		}
	}
	json.NewEncoder(w).Encode(siparisler)
}
