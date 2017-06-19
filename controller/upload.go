package controller

import (
	"github.com/labstack/echo"
	"os"
	"io"
	"net/http"
	"fmt"
)

var filepath string

func init() {
	filepath = "files/"
}

func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(filepath + file.Filename)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
}
