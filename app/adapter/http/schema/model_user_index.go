// DO NOT EDIT! Code generated by openapi-generator
package schema

// UserIndex - ユーザー情報（ユーザー一覧で使用）
type UserIndex struct {

	UserId int32 `json:"userId"`

	Nickname string `json:"nickname"`

	ImageUrl *string `json:"imageUrl"`

	Occupations []OccupationSummary `json:"occupations,omitempty"`

	ReviewScore float64 `json:"reviewScore"`

	Areas []AreaForUser `json:"areas"`

	Menu *MenuSummaryForUser `json:"menu"`

	IsFavorite bool `json:"isFavorite"`

	// 得意な施術リスト
	OccupationTags []string `json:"occupationTags"`

	ShortestReservableDate *string `json:"shortestReservableDate"`

	Menus []MenuForUserIndex `json:"menus"`
}
