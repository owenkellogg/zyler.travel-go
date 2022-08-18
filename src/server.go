package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/spf13/viper"
)

type Media struct {
	Url  string `json:"url" xml:"url"`
	Name string `json:"name" xml:"name"`
}

type MediaResponse struct {
	Item Media `json:"item" xml:"item"`
}

func upload(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	timestamp := time.Now().Unix()

	directory := viper.GetString("uploads_directory")

	fmt.Println(directory)

	filename := fmt.Sprintf("%s/%d-%s", directory, timestamp, file.Filename)

	// Destination
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	media := &Media{
		Url:  fmt.Sprintf("https://zyler.travel/public/uploads/%d-%s/", timestamp, file.Filename),
		Name: file.Filename,
	}

	return c.JSON(http.StatusOK, &MediaResponse{Item: *media})

}

func main() {

	fmt.Sprintln("START")

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	viper.BindEnv("uploads_directory")

	viper.SetDefault("port", ":5200")
	viper.SetDefault("uploads_directory", "public/uploads")

	port := viper.GetString("port")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}))

	e.Static("/", "public")

	e.POST("/api/upload", upload)

	e.Logger.Fatal(e.Start(port))
}
