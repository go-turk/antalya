package order

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

const NegativePriceError = "siparisin fiyatı negatif olamaz"

var Siparisler map[string]*Siparis

type Siparis struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	IsDelivered bool    `json:"is_delivered"`
	Price       float64 `json:"price"`
}

func init() {
	fmt.Println("Sipariş Kütüphanesi çalıştırıldı")
	Siparisler = map[string]*Siparis{}
}

func NewSiparis(description string, price float64) (*Siparis, error) {
	if price <= 0 {
		return nil, errors.New(NegativePriceError)
	}
	siparis := &Siparis{
		Code:        uuid.New().String(),
		Description: description,
		IsDelivered: false,
		Price:       price,
	}
	Siparisler[siparis.Code] = siparis
	return siparis, nil
}
