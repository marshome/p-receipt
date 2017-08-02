package models

type ReceiptInfo struct {
	Lang       string  `json:"lang,omitempty"`
	Title      string  `json:"title,omitempty"`
	TotalPrice float64 `json:"totalPrice,omitempty"`
	FullText   string  `json:"fullText,omitempty"`
}
