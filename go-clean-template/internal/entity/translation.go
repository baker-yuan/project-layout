// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation 翻译记录
type Translation struct {
	Source      string `json:"source"       example:"auto"`                 // 输入语言
	Destination string `json:"destination"  example:"en"`                   // 返回语言
	Original    string `json:"original"     example:"текст для перевода"`   // 输入内容
	Translation string `json:"translation"  example:"text for translation"` // 翻译结果
}
