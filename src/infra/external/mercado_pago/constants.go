package mercado_pago

func GetMercadoPagoPostUrl() string {
	return "https://api.mercadopago.com/instore/orders/qr/seller/collectors/214927776/pos/SUC001POS001/qrs"
}

func GetAppNotificationUrl() string {
	return "https://www.yourserver.com/notifications"
}

func GetMarcadoPagoToken() string {
	return "TEST-3413529352651645-022515-2fd495e4a5258bb49c028d9893446377-214927776"
}

func GetBearerToken() string {
	return "Bearer " + GetMarcadoPagoToken()
}
