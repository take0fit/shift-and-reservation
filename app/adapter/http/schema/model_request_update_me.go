// DO NOT EDIT! Code generated by openapi-generator
package schema

type RequestUpdateMe struct {

	// ニックネーム
	Nickname string `json:"nickname"`

	// 氏名（姓）
	LastName string `json:"lastName"`

	// 氏名（名）
	FirstName string `json:"firstName"`

	// カナ（姓）
	LastNameRuby string `json:"lastNameRuby"`

	// カナ（名）
	FirstNameRuby string `json:"firstNameRuby"`

	// 性別（0:男 1:女 2:その他）
	Sex int32 `json:"sex"`

	// 生年月日
	Birthday string `json:"birthday"`

	// 郵便番号
	ZipCd string `json:"zipCd"`

	// 都道府県ID
	PrefectureId int32 `json:"prefectureId"`

	// 市区町村
	City string `json:"city"`

	// 携帯電話
	Phone string `json:"phone"`
}
