package mercado_pago

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"io"
	"net/http"
)

type MercadoPagoIntegration struct {
}

func (m MercadoPagoIntegration) CreatePayment(order entities.Order) (string, error) {
	var qrCode string

	orderJson, err := json.Marshal(MountMercadoPagoRequest(order))
	if err != nil {
		return qrCode, errors.New("erro ao deserializar o objeto")
	}

	var bearer = GetBearerToken()

	r, err := http.NewRequest("POST", GetMercadoPagoPostUrl(), bytes.NewBuffer(orderJson))
	if err != nil {
		return qrCode, errors.New("erro ao montar o request ao mercado pago")
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return qrCode, errors.New("erro ao fazer o post ao mercado pago")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)

	var apiResponse *QrCreatedResponse

	derr := json.NewDecoder(res.Body).Decode(apiResponse)
	if derr != nil {
		return qrCode, errors.New("erro ao deserializar a resposta do mercado pago")
	}

	if res.StatusCode != http.StatusCreated {
		return qrCode, errors.New("mercado pago retornou um erro")
	}

	return apiResponse.QrData, nil
}
