package util

import (
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func EchoRenderTempl(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}


func MakeDirectory(path string) error {

    _, err := os.Stat(path)

	if os.IsNotExist(err) {

		if err = os.MkdirAll(path, os.ModePerm); err != nil {

            return err
		}

		_, err = os.Stat(path)
	}

	if err != nil {
        return err
	}

    return nil
}

