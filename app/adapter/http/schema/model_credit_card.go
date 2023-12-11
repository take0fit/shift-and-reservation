// DO NOT EDIT! Code generated by openapi-generator
package schema

// CreditCard - クレジットカード情報
type CreditCard struct {

	// クレジットカードID
	CreditCardId string `json:"creditCardId"`

	// ブランド名
	Brand string `json:"brand"`

	// カード番号末尾4桁
	CardNumberLast4 string `json:"cardNumberLast4"`

	// カード有効期限（年）
	ExpYear int32 `json:"expYear"`

	// カード有効期限（月）
	ExpMonth int32 `json:"expMonth"`

	// カード名義人
	Name string `json:"name"`

	// デフォルトカードフラグ
	IsDefaultCard bool `json:"isDefaultCard"`
}
