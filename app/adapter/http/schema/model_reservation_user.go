// DO NOT EDIT! Code generated by openapi-generator
package schema

type ReservationUser struct {

	// ビューティシャンID
	UserId int32 `json:"userId"`

	// ニックネーム
	Nickname string `json:"nickname"`

	// プロフィール画像URL
	ProfileImageUrl *string `json:"profileImageUrl"`

	// 注意事項
	Cautions string `json:"cautions"`
}