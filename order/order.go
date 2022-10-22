package order

import (
	"fmt"
	"github.com/google/uuid"
)

var Siparisler map[*Siparis]bool

type Siparis struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	IsDelivered bool   `json:"is_delivered"`
}

func init() {
	fmt.Println("Sipariş Kütüphanesi çalıştırıldı")
	Siparisler = map[*Siparis]bool{}
}

func NewSiparis(description string) *Siparis {
	siparis := &Siparis{
		Code:        uuid.New().String(),
		Description: description,
		IsDelivered: false,
	}
	Siparisler[siparis] = true
	return siparis
}
