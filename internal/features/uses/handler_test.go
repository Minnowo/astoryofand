package uses

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
    passingUsesJSON = []string{
        `{
        "email": "test@gmail.com",
        "fullname": "test",
        "comment": "test"
        }`,
	}

	failingUsesJSON = []string{
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

func TestUsesPlaced(t *testing.T) {

	e := echo.New()

	h := &UsesHandler{
		EncryptionWriter: &crypto.PGPEncryptionWriter{
			PublicKey:       assets.PublicKeyBytes,
			OutputDirectory: t.TempDir(),
		},
	}

	// check valid orders
	for _, order := range passingUsesJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, h.HandleUsesPOST(c)) {
			assert.Equal(t, http.StatusPermanentRedirect, rec.Code)
		}
	}

	// check for bad requests to make sure they don't work
	for _, order := range failingUsesJSON {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(order))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		assert.Error(t, h.HandleUsesPOST(c))
	}
}
