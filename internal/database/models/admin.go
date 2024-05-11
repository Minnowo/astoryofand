package models

type AdminView struct {
	BoxSetPrice float32 `form:"boxpricevalue"`
	StickerCost float32 `form:"stickerpricevalue"`
	SignPGPKey  string
}
