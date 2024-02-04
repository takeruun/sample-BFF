package main

import "github.com/labstack/echo"

func main() {
	echo.New().Start(":3001")
}
