package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadv184/gopayment"
	"github.com/mohammadv184/gopayment/gateway/idpay"
	"github.com/mohammadv184/gopayment/gateway/payping"
	"log"
	"net/http"
)

func setPageAndData(c *gin.Context, data gin.H) {
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		data,
	)
}

// IndexHandler serves the index.html page
func IndexHandler(c *gin.Context) {
	log.Println("Loading main page")
	setPageAndData(c, gin.H{
		"page":    "main",
		"drivers": Drivers,
	})
}

// PreviewHandler serves the preview.html page
func PreviewHandler(c *gin.Context) {
	log.Println("Loading preview page")
	setPageAndData(c, gin.H{
		"page":   "preview",
		"driver": c.Param("driver"),
	})
}

// PaymentHandler serves the payment.html page
func PaymentHandler(c *gin.Context) {
	log.Println("Loading payment page")
	payment := gopayment.NewPayment(Drivers[c.Param("driver")])
	payment.Amount(1000)
	err := payment.Purchase()
	if err != nil {
		setPageAndData(c, gin.H{
			"page":   "result",
			"status": "error",
			"msg":    err.Error(),
			"img":    "failed",
		})
		return
	}

	setPageAndData(c, gin.H{
		"page":   "payment",
		"method": payment.PayMethod(),
		"payURL": payment.PayURL(),
	})
}

// CallBackHandler serves the result.html page
func CallBackHandler(c *gin.Context) {
	log.Println("Loading result page")
	gateway := Drivers[c.Param("driver")]
	switch c.Param("driver") {
	case payping.Driver{}.GetDriverName():
		vReq := &payping.VerifyRequest{
			RefID:  c.PostForm("refid"),
			Amount: "1000",
		}
		receipt, err := gateway.Verify(vReq)
		if err != nil {
			setPageAndData(c, gin.H{
				"page":   "result",
				"status": "failed",
				"msg":    err.Error(),
				"img":    "failed",
			})
			return
		}
		setPageAndData(c, gin.H{
			"page":   "result",
			"status": "success",
			"msg": "Payment is successfully refrenceCode: " +
				receipt.GetReferenceID() + "\n Driver is: " +
				receipt.GetDriver() + "\n Date:" + receipt.GetDate().Format("2006-01-02 15:04:05"),
			"img": "success",
		})
	case idpay.Driver{}.GetDriverName():
		vReq := &idpay.VerifyRequest{
			RefID: c.PostForm("order_id"),
			ID:    c.PostForm("id"),
		}
		receipt, err := gateway.Verify(vReq)
		if err != nil {
			setPageAndData(c, gin.H{
				"page":   "result",
				"status": "failed",
				"msg":    err.Error(),
				"img":    "failed",
			})
			return
		}

		setPageAndData(c, gin.H{
			"page":   "result",
			"status": "success",
			"msg": "Payment is successfully refrenceCode: " +
				receipt.GetReferenceID() + "\n Driver is: " +
				receipt.GetDriver() + "\n Date:" + receipt.GetDate().Format("2006-01-02 15:04:05") +
				"\n Card Number: " + receipt.GetDetail("cardNumber") + "\n Hashed Card Number: " + receipt.GetDetail("HashedCardNumber"),
			"img": "success",
		})

	default:
		log.Println("Driver not found")

	}

}
