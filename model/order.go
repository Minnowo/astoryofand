package model

type Order struct {
	Email          string  `json:"email" form:"email"`
	PayMethod      string  `json:"paymethod" form:"paymethod"`
	BoxSetCount    uint32  `json:"boxsetcount" form:"boxsetcount"`
	StickerCount   uint32  `json:"stickercount" form:"stickercount"`
	TotalCost      float32 `json:"totalcost"`
	FullName       string  `json:"fullname" form:"fullname"`
	DeliveryMethod string  `json:"deliverymethod" form:"deliverymethod"`
	Address        string  `json:"address" form:"address"`
	City           string  `json:"city" form:"city"`
	ZipCode        string  `json:"zipcode" form:"zipcode"`
	OtherDelivery  string  `json:"otherdelivery" form:"otherdelivery"`
	OtherPay       string  `json:"otherpay" form:"otherpay"`
}

func (o *Order) CheckValidOrder() bool {

	if IsEmptyOrWhitespace(o.Email) ||
		IsEmptyOrWhitespace(o.PayMethod) ||
		IsEmptyOrWhitespace(o.FullName) ||
		IsEmptyOrWhitespace(o.DeliveryMethod) {
		return false
	}

	if o.BoxSetCount < 1 || o.StickerCount < 0 {
		return false
	}

	return true
}
