package order

import (
	"fmt"
	"github.com/google/uuid"
)

var Siparisler map[string]*Siparis

type Siparis struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	IsDelivered bool   `json:"is_delivered"`
}

func init() {
	fmt.Println("Sipariş Kütüphanesi çalıştırıldı")
	Siparisler = map[string]*Siparis{}
}

func NewSiparis(description string) *Siparis {
	siparis := &Siparis{
		Code:        uuid.New().String(),
		Description: description,
		IsDelivered: false,
	}
	Siparisler[siparis.Code] = siparis
	return siparis
}
