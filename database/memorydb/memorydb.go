package memorydb

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/assets"
)

type MemoryDB struct {
	boxSetPrice float32
	stickerCost float32

	sync.RWMutex
}

var db MemoryDB

func InitDB() {

	db = MemoryDB{
		boxSetPrice: assets.BoxSetPrice,
		stickerCost: assets.StickerCost,
	}
}

func GetDB() *MemoryDB {
	return &db
}

func (d *MemoryDB) GetBoxPrice() float32 {
	d.RLock()
	r := d.boxSetPrice
	d.RUnlock()
	return r
}

func (d *MemoryDB) SetBoxPrice(newPrice float32) {
	if newPrice <= 0 {
		return
	}

	d.Lock()
	d.boxSetPrice = newPrice
	d.Unlock()

	log.Infof("Box price has changed to %f", newPrice)
}

func (d *MemoryDB) GetStickerPrice() float32 {
	d.RLock()
	r := d.stickerCost
	d.RUnlock()
	return r
}

func (d *MemoryDB) SetStickerPrice(newPrice float32) {
	if newPrice <= 0 {
		return
	}

	d.Lock()
	d.stickerCost = newPrice
	d.Unlock()

	log.Infof("Sticker price has changed to %f", newPrice)
}
