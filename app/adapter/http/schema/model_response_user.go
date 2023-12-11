// DO NOT EDIT! Code generated by openapi-generator
package schema

// ResponseUser - ユーザー詳細
type ResponseUser struct {

	// ユーザーID
	UserId int32 `json:"userId"`

	// ニックネーム
	Nickname string `json:"nickname"`

	// プロフィール画像
	ProfileImageUrl *string `json:"profileImageUrl"`

	// 自己紹介
	SelfIntroduction string `json:"selfIntroduction"`

	// 性別（0:男 1:女 2:その他）
	Sex int32 `json:"sex"`
}
