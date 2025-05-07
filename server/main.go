package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{
			"https://m8-final-bucket.s3.us-east-1.amazonaws.com/13439846_1920_1080_25fps.mp4",
			"https://m8-final-bucket.s3.us-east-1.amazonaws.com/13432587_3840_2160_30fps.mp4",
			"https://m8-final-bucket.s3.us-east-1.amazonaws.com/12981408_3840_2160_25fps.mp4",
			"https://m8-final-bucket.s3.us-east-1.amazonaws.com/13489700_3840_2160_30fps.mp4",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}