package mercado_pago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"io"
	"net/http"
)

func PostPayment(order entities.Order) {
	orderJson, err := json.Marshal(MountMercadoPagoRequest(order))
	if err != nil {
		fmt.Println(err)
		return
	}

	var bearer = GetBearerToken()

	r, err := http.NewRequest("POST", GetMercadoPagoPostUrl(), bytes.NewBuffer(orderJson))
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

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}
}
