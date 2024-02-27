package mercado_pago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Sponsor struct {
	Id int `json:"id"`
}

type Item struct {
	SkuNumber   string `json:"sku_number"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnitPrice   int    `json:"unit_price"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure"`
	TotalAmount int    `json:"total_amount"`
}

type CreateQR struct {
	Description       string    `json:"description"`
	ExpirationDate    time.Time `json:"expiration_date"`
	ExternalReference string    `json:"external_reference"`
	Items             []Item    `json:"items"`
	NotificationUrl   string    `json:"notification_url"`
	Sponsor           Sponsor   `json:"sponsor"`
	Title             string    `json:"title"`
	TotalAmount       int       `json:"total_amount"`
}

type QrCreatedResponse struct {
	InStoreOrderId string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

func PostPayment() {
	posturl := "https://api.mercadopago.com/instore/orders/qr/seller/collectors/214927776/pos/SUC001POS001/qrs"

	body := []byte(`{
    "description": "Purchase description.",
    "expiration_date": "2024-02-26T16:34:56.559-04:00",
    "external_reference": "[id_do_pedido]",
    "items": [
        {
            "sku_number": "A123K9191938",
            "category": "marketplace",
            "title": "Point Mini",
            "description": "This is the Point Mini",
            "unit_price": 100,
            "quantity": 1,
            "unit_measure": "unit",
            "total_amount": 100
        }
    ],
    "notification_url": "https://www.yourserver.com/notifications",
    "sponsor": {
        "id": 262583389
    },
    "title": "Pedido #[id_do_pedido]",
    "total_amount": 100
}`)

	var bearer = "Bearer " + "TEST-3413529352651645-022515-2fd495e4a5258bb49c028d9893446377-214927776"

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	post := &QrCreatedResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	fmt.Println(post.InStoreOrderId)
	fmt.Println(post.QrData)

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}
}
