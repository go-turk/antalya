package order

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

func SiparisVer(w http.ResponseWriter, r *http.Request) {
	// bir dışarıdan gelen veriyi tutacağımız nesneyi oluşturduk
	var requestBody struct {
		Description string `json:"description"`
		IsUser      bool   `json:"is_user"`
		IsUserID    int    `json:"is_user_id"`
	}
	// dışarıdan gelen veriyi tuttuk
	json.NewDecoder(r.Body).Decode(&requestBody)

	if requestBody.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Lütfen sipariş açıklaması giriniz."))
		return
	}

	if !requestBody.IsUser {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sadece Bu hizmetten müşterilerimiz yararlanabilir. Lütfen öncelikli olarak kayıt olunuz."))
		return
	}

	if requestBody.IsUserID%7 != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Böyle bir kullanıcı bulunamadı! Lütfen UserID'nizi doğru girdiğinden emin olunuz!"))
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

func TeslimEt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	siparis := Siparisler[id]

	//Sipariş teslim edilmeden iptal edilmişse engellenir.
	if siparis.IsCanceled == true {
		w.Write([]byte("İade edilmiş bir siparişi teslim edemezsin..."))
		return
		//Teslim edilen ürün engelelnir.
	} else if siparis.IsDelivered == true {
		w.Write([]byte("Ürün zaten teslim edilmiş..."))
		return
		//Ürün teslim edilir.
	} else if siparis != nil {
		siparis.IsDelivered = true
		w.Write([]byte("Siparişiniz teslim edildi..."))
		return
	}
	w.Write([]byte("Sipariş Bulunamadı..."))
}

func IptalEt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	siparis := Siparisler[id]

	//Ürün iptal edilmişse ve teslim edilmemişse iptal edilemez.
	if siparis.IsCanceled == true && siparis.IsDelivered == false {
		w.Write([]byte("Ürün zaten iptal edilmiş..."))
		return
		//Ürün teslim edilmemişse iptal edilir.
	} else if siparis.IsDelivered == false {
		siparis.IsCanceled = true
		w.Write([]byte("Sipariş İptal Edildi..."))
		return
		//Teslim edilmişse iade süreci başlar
	} else if siparis.IsDelivered == true {
		siparis.IsDelivered = false
		w.Write([]byte("İade süreci başlatıldı.."))
		siparis.IsCanceled = true
		w.Write([]byte("iade edildi..."))
		return
	}

	//Bu kurguyu çalıştıramadım.
	w.Write([]byte("Sipariş Bulunamadı..."))

}

func FiyatSorgula(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		StoreLocation struct {
		X float64
		Y float64
		} `json:"store_location"`
		CustomerLocation struct {
		X float64
		Y float64
		} `json:"customer_location"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
	}

	width := requestBody.StoreLocation.X - requestBody.CustomerLocation.X
	height := requestBody.StoreLocation.Y - requestBody.CustomerLocation.Y
	hypotenuse := math.Sqrt(width*width + height*height)

	price := hypotenuse * 25

	response := map[string]interface{}{
		"message": "price calculated successfully!",
		"price": price,
	}


	json.NewEncoder(w).Encode(response)	
}