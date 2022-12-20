package web

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammadv184/gopayment/gateway"
	"github.com/mohammadv184/gopayment/gateway/asanpardakht"
	"github.com/mohammadv184/gopayment/gateway/idpay"
	"github.com/mohammadv184/gopayment/gateway/payping"
	"github.com/mohammadv184/gopayment/gateway/zarinpal"
	"github.com/mohammadv184/gopayment/gateway/zibal"
)

var Drivers map[string]gateway.Driver

func Init() {
	_ = godotenv.Load("./.env")
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
	router.Any("/callback/:driver", CallBackHandler)

	// Start and run the server
	err := router.Run(":3000")
	if err != nil {
		log.Println("Error: ", err)
	}
}

func registerDrivers() {
	Drivers = make(map[string]gateway.Driver)
	Drivers[payping.Driver{}.GetDriverName()] = &payping.Driver{
		Token:    os.Getenv("PAYPING_TOKEN"),
		Callback: os.Getenv("PAYPING_CALLBACK"),
	}
	Drivers[zarinpal.Driver{}.GetDriverName()] = &zarinpal.Driver{
		MerchantID: os.Getenv("ZARINPAL_MERCHANT_ID"),
		Callback:   os.Getenv("ZARINPAL_CALLBACK"),
	}
	Drivers[idpay.Driver{}.GetDriverName()] = &idpay.Driver{
		MerchantID: os.Getenv("IDPAY_MERCHANT_ID"),
		Callback:   os.Getenv("IDPAY_CALLBACK"),
		Sandbox:    os.Getenv("IDPAY_SANDBOX") == "true",
	}
	Drivers[asanpardakht.Driver{}.GetDriverName()] = &asanpardakht.Driver{
		MerchantConfigID: os.Getenv("ASANPARDAKHT_MERCHANT_CONFIG_ID"),
		Callback:         os.Getenv("ASANPARDAKHT_CALLBACK"),
		Username:         os.Getenv("ASANPARDAKHT_USERNAME"),
		Password:         os.Getenv("ASANPARDAKHT_PASSWORD"),
	}
	Drivers[zibal.Driver{}.GetDriverName()] = &zibal.Driver{
		Merchant: os.Getenv("ZIBAL_MERCHANT"),
		Callback: os.Getenv("ZIBAL_CALLBACK"),
	}
}
