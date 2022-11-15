package dto

import "time"

type MProduct struct {
	ID                        uint
	ProCode                   string        `json:"procode"`
	Sku                       string        `json:"sku"`
	ProName                   string        `json:"proname"`
	ProName2                  string        `json:"pronam2"`
	GenerikName               string        `json:"generik_name"`
	Kategori                  int           `json:"generik_name"`
	Etalase                   string        `json:"etalase"`
	FotoUtama                 string        `json:"foto_utama"`
	FotoDepan                 string        `json:"foto_depan"`
	FotoSamping               string        `json:"foto_samping"`
	FotoAtas                  string        `json:"foto_atas"`
	FotoDetail                string        `json:"foto_detail"`
	Deskripsi                 []interface{} `json:"deskripsi"`
	LinkVideo                 string        `json:"link_video"`
	Custom                    string        `json:"custom"`
	SellUnit                  int           `json:"sell_unit"`
	SellPack                  int           `json:"sell_pack"`
	SellPackName              string        `json:"sell_pack_name"`
	CustomSellUnit            int           `json:"custom_sell_unit"`
	CustomSellPack            int           `json:"custom_sell_pack"`
	CustomSellPackName        string        `json:"custom_sell_pack_name"`
	Berat                     int           `json:"berat"`
	HargaMin                  uint          `json:"harga_min"`
	HargaMax                  uint          `json:"harga_max"`
	HargaPdn                  uint          `json:"harga_pdn"`
	HargaIndoApotek           uint          `json:"harga_indo_apotek"`
	Wajib                     string        `json:"wajib"`
	ApotekOnline              string        `json:"apotek_online"`
	IndoApotek                string        `json:"indo_apotek"`
	Green                     string        `json:"green"`
	UserId                    string        `json:"user_id"`
	LastUpdate                *time.Time    `json:"last_update"`
	HargaIndoApotekLastUpdate *time.Time    `json:"harga_indo_apotek_last_update"`
}
