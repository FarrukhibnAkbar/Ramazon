package models

type Ramazon struct {
	Bomdod     string `json:"bomdod" gorm:"bomdod"`
	Peshin     string `json:"peshin" gorm:"peshin"`
	Asr        string `json:"asr" gorm:"asr"`
	Shom       string `json:"shom" gorm:"shom"`
	Xufton     string `json:"xufton" gorm:"xufton"`
	Date       string `json:"date" gorm:"date"`
	HijriYear  string `json:"hijri_year" gorm:"hijri_year"`
	HijriMonth string `json:"hijri_month" gorm:"hijri_month"`
	HijriDay   string `json:"hijri_day" gorm:"hijri_day"`
}
