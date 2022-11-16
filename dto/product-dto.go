package dto

import (
	"gorm.io/datatypes"
	"time"
)

// table product a
type MProduct struct {
	ID                        uint
	ProCode                   string          `json:"procode" gorm:"type:varchar(256)"`
	Sku                       string          `json:"sku"`
	ProName                   string          `json:"proname" gorm:"type:varchar(256)"`
	ProName2                  string          `json:"pronam2" gorm:"type:varchar(256)"`
	GenerikName               string          `json:"generik_name" gorm:"type:varchar(256)"`
	Kategori                  int             `json:"generik_name" gorm:"type:varchar(256)"`
	Etalase                   string          `json:"etalase" gorm:"type:varchar(256)"`
	FotoUtama                 string          `json:"foto_utama" gorm:"type:varchar(256)"`
	FotoDepan                 string          `json:"foto_depan" gorm:"type:varchar(256)"`
	FotoSamping               string          `json:"foto_samping" gorm:"type:varchar(256)"`
	FotoAtas                  string          `json:"foto_atas" gorm:"type:varchar(256)"`
	FotoDetail                string          `json:"foto_detail" gorm:"type:varchar(256)"`
	Deskripsi                 *datatypes.JSON `json:"deskripsi"`
	LinkVideo                 string          `json:"link_video"`
	Custom                    string          `json:"custom" gorm:"type:varchar(2)"`
	SellUnit                  int             `json:"sell_unit"`
	SellPack                  int             `json:"sell_pack"`
	SellPackName              string          `json:"sell_pack_name" gorm:"type:varchar(256)"`
	CustomSellUnit            int             `json:"custom_sell_unit"`
	CustomSellPack            int             `json:"custom_sell_pack"`
	CustomSellPackName        string          `json:"custom_sell_pack_name" gorm:"type:varchar(256)"`
	Berat                     int             `json:"berat"`
	HargaMin                  uint            `json:"harga_min"`
	HargaMax                  uint            `json:"harga_max"`
	HargaPdn                  uint            `json:"harga_pdn"`
	HargaIndoApotek           uint            `json:"harga_indo_apotek"`
	Wajib                     string          `json:"wajib" gorm:"type:varchar(2)"`
	ApotekOnline              string          `json:"apotek_online" gorm:"type:varchar(2)"`
	IndoApotek                string          `json:"indo_apotek" gorm:"type:varchar(2)"`
	Green                     string          `json:"green" gorm:"type:varchar(2)"`
	UserId                    string          `json:"user_id" gorm:"type:varchar(256)"`
	LastUpdate                *time.Time      `json:"last_update"`
	HargaIndoApotekLastUpdate *time.Time      `json:"harga_indo_apotek_last_update"`
}

type MProductRequest struct {
	Data []ProductRequestMaster
}

type ProductRequestMaster struct {
	ID                        uint
	ProCode                   string        `json:"procode" gorm:"type:varchar(256)"`
	Sku                       string        `json:"sku"`
	ProName                   string        `json:"proname" gorm:"type:varchar(256)"`
	ProName2                  string        `json:"pronam2" gorm:"type:varchar(256)"`
	GenerikName               string        `json:"generik_name" gorm:"type:varchar(256)"`
	Kategori                  int           `json:"generik_name" gorm:"type:varchar(256)"`
	Etalase                   string        `json:"etalase" gorm:"type:varchar(256)"`
	FotoUtama                 string        `json:"foto_utama" gorm:"type:varchar(256)"`
	FotoDepan                 string        `json:"foto_depan" gorm:"type:varchar(256)"`
	FotoSamping               string        `json:"foto_samping" gorm:"type:varchar(256)"`
	FotoAtas                  string        `json:"foto_atas" gorm:"type:varchar(256)"`
	FotoDetail                string        `json:"foto_detail" gorm:"type:varchar(256)"`
	Deskripsi                 []interface{} `json:"deskripsi"`
	LinkVideo                 string        `json:"link_video"`
	Custom                    string        `json:"custom" gorm:"type:varchar(2)"`
	SellUnit                  int           `json:"sell_unit"`
	SellPack                  int           `json:"sell_pack"`
	SellPackName              string        `json:"sell_pack_name" gorm:"type:varchar(256)"`
	CustomSellUnit            int           `json:"custom_sell_unit"`
	CustomSellPack            int           `json:"custom_sell_pack"`
	CustomSellPackName        string        `json:"custom_sell_pack_name" gorm:"type:varchar(256)"`
	Berat                     int           `json:"berat"`
	HargaMin                  uint          `json:"hargamin"`
	HargaMax                  uint          `json:"hargamax"`
	HargaPdn                  uint          `json:"harga_pdn"`
	HargaIndoApotek           uint          `json:"harga_indoapotek"`
	Wajib                     string        `json:"wajib" gorm:"type:varchar(2)"`
	ApotekOnline              string        `json:"apotek_online" gorm:"type:varchar(2)"`
	IndoApotek                string        `json:"indo_apotek" gorm:"type:varchar(2)"`
	Green                     string        `json:"green" gorm :"type:varchar(2)"`
	UserId                    string        `json:"user_id" gorm:"type:varchar(256)"`
	LastUpdate                time.Time     `json:"last_update"`
	HargaIndoApotekLastUpdate time.Time     `json:"harga_indo_apotek_last_update"`
}
