// DO NOT EDIT! Code generated by openapi-generator
package schema

// ShiftPerMonth - 一月の日毎のシフト情報
type ShiftPerMonth struct {

	// 日付
	Date string `json:"date"`

	// エリア名
	AreaName string `json:"areaName"`

	AvailableHours []string `json:"availableHours"`
}
