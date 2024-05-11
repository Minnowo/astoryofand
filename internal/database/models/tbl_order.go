package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

type TableOrder struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UUID      string         `gorm:"primarykey" json:"uuid"`
	UserData
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

func NewOrder() *TableOrder {
	return &TableOrder{
		UserData: UserData{OrderType},
	}
}

func (o *TableOrder) DelayedInit() *TableOrder {
	if err := uuid.Validate(o.UUID); err != nil {

		o.UUID = util.GetOrderID()
	}

	o.CreatedAt = time.Now()
	o.UpdatedAt = o.CreatedAt

	return o
}

func (o *TableOrder) EnsureType() *TableOrder {
	o.Type = OrderType
	return o
}

func (o *TableOrder) CheckValidDataFromUser() error {

	if util.IsEmptyOrWhitespace(o.Email) ||
		util.IsEmptyOrWhitespace(o.PayMethod) ||
		util.IsEmptyOrWhitespace(o.FullName) ||
		util.IsEmptyOrWhitespace(o.DeliveryMethod) {
		return fmt.Errorf("Missing required string values! (Email, Paymethod, Fullname, Delivery, etc)")
	}

	if o.BoxSetCount < 0 || o.StickerCount < 0 {
		return fmt.Errorf("Box set or sticker count must be > 0!")
	}

	return nil
}
