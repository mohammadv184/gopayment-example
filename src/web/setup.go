package web

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammadv184/gopayment/gateway"
	"github.com/mohammadv184/gopayment/gateway/idpay"
	"github.com/mohammadv184/gopayment/gateway/payping"
	"github.com/mohammadv184/gopayment/gateway/zarinpal"
	"os"
)

var Drivers map[string]gateway.Driver

func Init() {
	godotenv.Load("./.env")

	registerDrivers()
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	// Serve HTML templates
	router.LoadHTMLGlob("./templates/*")
	// Serve frontend static files
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))

	// setup client side templates
	router.GET("/", IndexHandler)
	router.GET("/preview/:driver", PreviewHandler)
	router.GET("/payment/:driver", PaymentHandler)
	router.POST("/callback/:driver", CallBackHandler)

	// Start and run the server
	router.Run(":3000")
}

func registerDrivers() {
	Drivers = make(map[string]gateway.Driver)
	Drivers[payping.Driver{}.GetDriverName()] = &payping.Driver{
		Token:       os.Getenv("PAYPING_TOKEN"),
		Description: os.Getenv("PAYPING_DESCRIPTION"),
		Callback:    os.Getenv("PAYPING_CALLBACK"),
	}
	Drivers[zarinpal.Driver{}.GetDriverName()] = &zarinpal.Driver{
		MerchantID:  os.Getenv("ZARINPAL_MERCHANT_ID"),
		Description: os.Getenv("ZARINPAL_DESCRIPTION"),
		Callback:    os.Getenv("ZARINPAL_CALLBACK"),
	}
	Drivers[idpay.Driver{}.GetDriverName()] = &idpay.Driver{
		MerchantID:  os.Getenv("IDPAY_MERCHANT_ID"),
		Description: os.Getenv("IDPAY_DESCRIPTION"),
		Callback:    os.Getenv("IDPAY_CALLBACK"),
		Sandbox:     os.Getenv("IDPAY_SANDBOX") == "true",
	}
}
