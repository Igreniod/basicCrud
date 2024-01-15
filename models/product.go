package models

type Product struct {
	Id          int64   `gorm:"primarykey" json:"id"`
	NamaProduct string  `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string  `gorm:"type:text" json:"deskripsi"`
	Harga       float64 `gorm:"type:decimal(10,2)" json:"harga"`
	Diskon      int     `gorm:"type:int" json:"diskon"`
	HargaDiskon float64 `gorm:"type:decimal(10,2)" json:"harga_diskon"`
}
