package contact

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/crypto"
	"github.com/stretchr/testify/assert"
)

var (
	passingJSON = []string{
		`{
        "email": "test@gmail.com",
        "fullname": "test",
        "comment": "test"
        }`,
	}

	failingJSON = []string{
		// email is empty
		`{
        "email": "",
        "fullname": "test",
        "comment": "test"
        }`,
		// name is empty
		`{
        "email": "test@test.test",
        "fullname": "",
        "comment": "test"
        }`,
		// comment is empty
		`{
        "email": "test@test.test",
        "fullname": "test",
        "comment": ""
        }`,
	}
)

func TestOrderPlaced(t *testing.T) {

	e := echo.New()

	h := &ContactUsHandler{
		EncryptionWriter: &crypto.PGPEncryptionWriter{
			PublicKey:       assets.PublicKeyBytes,
			OutputDirectory: t.TempDir(),
		},
	}

	// check valid orders
	for _, order := range passingJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, h.POSTContactUsPlaced(c)) {
			assert.Equal(t, http.StatusSeeOther, rec.Code)
		}
	}

	// check for bad requests to make sure they don't work
	for _, order := range failingJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		assert.Error(t, h.POSTContactUsPlaced(c))
	}
}
