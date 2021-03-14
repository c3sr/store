package web

import (
	"net/http"

	"github.com/k0kubun/pp/v3"
	"github.com/labstack/echo/v4"
	"github.com/c3sr/store"
	"github.com/c3sr/uuid"
)

func put(c echo.Context) error {
	store := c.Get("store").(store.Store)
	file, err := c.FormFile("files")
	if err != nil {
		pp.Println("ERRROR..... " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer src.Close()

	id := uuid.NewV4()
	key, err := store.UploadFrom(src, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"id": key,
	})
}
