package order

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/crypto"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/stretchr/testify/assert"
)

var (
	passingOrderJSON = []string{
		fmt.Sprintf(`{
            "email": "test@test.test",
            "paymethod": "cash",
            "boxpricetimeofbuy": %f,
            "totalcost": %f,
            "stickerpricetimeofbuy": %f,
            "boxsetcount": 1,
            "stickercount": 0,
            "fullname": "test_user",
            "deliverymethod": "delivery",
            "address": "test st south",
            "city": "testcity",
            "zipcode": "K2J0JM",
            "otherdelivery": "",
            "otherpay": ""
            }`,
			database.GetBoxPrice(),
			database.GetBoxPrice(),
			database.GetStickerPrice()),
	}

	failingOrderJSON = []string{
		// price is invalid
		fmt.Sprintf(`{
        "email": "another@test.test",
        "paymethod": "credit",
        "boxpricetimeofbuy": %f,
        "stickerpricetimeofbuy": 3,
        "boxsetcount": 2,
        "stickercount": 5,
        "totalcost": 50,
        "fullname": "another_user",
        "deliverymethod": "pickup",
        "address": "123 Main St",
        "city": "anothercity",
        "zipcode": "M3H5Y8",
        "otherdelivery": "express",
        "otherpay": "paypal"
        }`, database.GetBoxPrice()-1),

		// email is empty
		fmt.Sprintf(`{
        "email": "",
        "paymethod": "credit",
        "boxpricetimeofbuy": %f,
        "stickerpricetimeofbuy": 3,
        "boxsetcount": 2,
        "stickercount": 5,
        "totalcost": 50,
        "fullname": "another_user",
        "deliverymethod": "pickup",
        "address": "123 Main St",
        "city": "anothercity",
        "zipcode": "M3H5Y8",
        "otherdelivery": "express",
        "otherpay": "paypal"
        }`, database.GetBoxPrice()-1),

		// not even json
		"this is not valid json!",
	}
)

func TestOrderPlaced(t *testing.T) {

	e := echo.New()

	h := &OrderHandler{
		EncryptionWriter: &crypto.PGPEncryptionWriter{
			PublicKey:       assets.PublicKeyBytes,
			OutputDirectory: t.TempDir(),
		},
	}

	// check valid orders
	for _, order := range passingOrderJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, h.HandleOrderPlaced(c)) {
			assert.Equal(t, http.StatusSeeOther, rec.Code)
		}
	}

	// check for bad requests to make sure they don't work
	for _, order := range failingOrderJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		assert.Error(t, h.HandleOrderPlaced(c))
	}
}
