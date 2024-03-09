package model

import (
	"fmt"

	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/util"
)

type Order struct {
	Email          string  `json:"email" form:"email"`
	PayMethod      string  `json:"paymethod" form:"paymethod"`
	BoxSetValue    float32 `json:"boxpricetimeofbuy" form:"boxpricevalue"`
	StickerValue   float32 `json:"stickerpricetimeofbuy" form:"stickerpricevalue"`
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

func (o *Order) CheckValid() error {

	if util.IsEmptyOrWhitespace(o.Email) ||
		util.IsEmptyOrWhitespace(o.PayMethod) ||
		util.IsEmptyOrWhitespace(o.FullName) ||
		util.IsEmptyOrWhitespace(o.DeliveryMethod) {
		return fmt.Errorf("Missing required string values! (Email, Paymethod, Fullname, Delivery, etc)")
	}

	if o.BoxSetCount < 0 || o.StickerCount < 0 {
		return fmt.Errorf("Box set or sticker count must be > 0!")
	}

	if !util.AlmostEqual32(o.BoxSetValue, database.GetBoxPrice()) {
		return fmt.Errorf("The value of the box set has changed!")
	}

	if !util.AlmostEqual32(o.StickerValue, database.GetStickerPrice()) {
		return fmt.Errorf("The value of the sticker has changed!")
	}

	return nil
}
