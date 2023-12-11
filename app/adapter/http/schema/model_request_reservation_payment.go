// DO NOT EDIT! Code generated by openapi-generator
package schema

// RequestReservationPayment - 予約した施術の決済
type RequestReservationPayment struct {

	// PAYJPのカードID（使用しない場合は空文字で）
	PayjpCardId string `json:"payjpCardId"`

	// 支払う合計金額
	PaymentAmount int32 `json:"paymentAmount"`

	// 決済で利用するポイント
	UsePoint int32 `json:"usePoint"`
}