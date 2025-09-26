package models

type Bioskop struct {
    ID     int64   `gorm:"primaryKey" json:"id"`
    Nama   string  `gorm:"size:300" json:"nama"`
    Lokasi string  `gorm:"size:300" json:"lokasi"`
    Rating float32 `json:"rating"`
}

