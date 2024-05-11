package database

import (
	"errors"
	"fmt"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/crypto"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

var (
	settingsLock = sync.RWMutex{}

	settingsInstance = models.TableSettings{
		Profile:     "default",
		BoxSetPrice: assets.BoxSetPrice,
		StickerCost: assets.StickerCost,
		PublicKey:   assets.PublicKeyBytes,
	}
)

func LoadSettings(profile string) {

	settingsLock.Lock()

	settingsInstance.Profile = profile

	err := GetDB().
		Where(&models.TableSettings{Profile: profile}).
		First(&settingsInstance).Error

	settingsLock.Unlock()

	if err == nil {

		log.Info("Loaded settings: ", settingsInstance)

		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {

		log.Fatal(err)

		return
	}

	log.Warnf("Could not find profile %s. Using default values", profile)

	if err = SaveSettings(); err != nil {

		log.Fatal(err)
	}

}

func saveSettingsNoLock() error {
	return GetDB().Save(&settingsInstance).Error
}

func SaveSettings() error {

	var err error

	settingsLock.RLock()

	err = saveSettingsNoLock()

	settingsLock.RUnlock()

	return err
}

func GetBoxPrice() float32 {

	settingsLock.RLock()

	r := settingsInstance.BoxSetPrice

	settingsLock.RUnlock()

	return r
}

func SetBoxPrice(newPrice float32) {

	if newPrice <= 0 {
		return
	}

	settingsLock.Lock()

	settingsInstance.BoxSetPrice = newPrice

	saveSettingsNoLock()

	settingsLock.Unlock()

	log.Infof("Box price has changed to %f", newPrice)
}

func GetStickerPrice() float32 {

	settingsLock.RLock()

	r := settingsInstance.StickerCost

	settingsLock.RUnlock()

	return r
}

func SetStickerPrice(newPrice float32) {

	if newPrice <= 0 {

		log.Error("Invalid price provided")

		return
	}

	settingsLock.Lock()

	settingsInstance.StickerCost = newPrice

	saveSettingsNoLock()

	settingsLock.Unlock()

	log.Infof("Sticker price has changed to %f", newPrice)
}

func GetPublicKey() string {

	settingsLock.RLock()

	r := settingsInstance.PublicKey

	settingsLock.RUnlock()

	return r
}

func SetPublicKey(nv string) error {

	if !crypto.IsPublicKey(nv) {

		log.Error("Cannot use new public key since it is invalid")

		return fmt.Errorf("Invalid public key given")
	}

	settingsLock.Lock()

	settingsInstance.PublicKey = nv

	saveSettingsNoLock()

	settingsLock.Unlock()

	log.Infof("Public Key has changed dto %s", nv)

	return nil
}

func OrderHasValidPrice(o *models.TableOrder) bool {

	settingsLock.RLock()
	defer settingsLock.RUnlock()

	if !util.AlmostEqual32(o.BoxSetValue, settingsInstance.BoxSetPrice) {
		return false
	}

	if !util.AlmostEqual32(o.StickerValue, settingsInstance.StickerCost) {
		return false
	}

	return true
}
