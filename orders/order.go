package orders

import "minno/astoryofand/assets"

type Order struct {
	Email          string `json:"email" form:"email"`
	PayMethod      string `json:"paymethod" form:"paymethod"`
	BoxSetCount    uint32 `json:"boxsetcount" form:"boxsetcount"`
	StickerCount   uint32 `json:"stickercount" form:"stickercount"`
	FullName       string `json:"fullname" form:"fullname"`
	DeliveryMethod string `json:"deliverymethod" form:"deliverymethod"`
}

func CheckValidOrder(o Order) bool {

	if assets.IsEmptyOrWhitespace(o.Email) ||
		assets.IsEmptyOrWhitespace(o.PayMethod) ||
		assets.IsEmptyOrWhitespace(o.FullName) ||
		assets.IsEmptyOrWhitespace(o.DeliveryMethod) {
		return false
	}

	if o.BoxSetCount < 1 || o.StickerCount < 0 {
		return false
	}

	return true
}
