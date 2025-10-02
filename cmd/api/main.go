package main

import (
	"context"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/modules"
)

// @title   Loyalty Backend
// @version 1.0

// @contact.name  CuongLV
// @contact.email cuonglv@smartosc.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	modules.NewApp().Run(context.Background())
}
