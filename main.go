package main

import (
	"fmt"
	"net/http"

	"github.com/go-turk/antalya/order"
	"github.com/gorilla/mux"
)

func main() {
	//ilksiparisdenemeleri()
	// Birinci : Bir Router Tanımlayalım
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/siparis/ver", order.SiparisVer).Methods("POST")
	r.HandleFunc("/siparisler", order.TumSiparisler).Methods("GET")
	r.HandleFunc("/siparisler/tamamlandi/{uuid}", order.Tamamlandi).Methods("PUT")
	// siparişi teslim edil olarak değiştirmesini istiyoruz.
	fmt.Println(":9096 çalışmaya başladı")
	http.ListenAndServe(":9096", r)
	// - Sipariş ekleme
	// - Sipariş teslim etme
	// 1. adımda inMemory (Projenin içinde bir değişkene kaydedeceğiz)
	// 2. adımda bunları postgreSQL denilen Database'e kaydedeceğiz.
}

func ilksiparisdenemeleri() {
	ilkSiparis, err := order.NewSiparis("2 ekmek 1 çay", 15.75)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("siparis id:", ilkSiparis.Code)
	fmt.Println("siparis açıklaması:", ilkSiparis.Description)
	if ilkSiparis.IsDelivered {
		fmt.Println("Siraiş Teslim edildi")
	} else {
		fmt.Println("Sipariş henüz Teslim edilmedi")
	}
	order.Siparisler[ilkSiparis.Code] = ilkSiparis
	fmt.Println(order.Siparisler)
}
