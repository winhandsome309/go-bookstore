package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	orderService "go-bookstore/internal/order/service"

	"strconv"

	"go-bookstore/pkg/config"

	shippingModel "go-bookstore/internal/shipping/model"

	"github.com/gin-gonic/gin"

	b64 "encoding/base64"

	"github.com/tidwall/gjson"
)

type Payload struct {
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	StoreId     string `json:"storeId"`
	RequestType string `json:"requestType"`
	IpnUrl      string `json:"ipnUrl"`
	RedirectUrl string `json:"redirectUrl"`
	OrderId     string `json:"orderId"`
	Amount      string `json:"amount"`
	Lang        string `json:"lang"`
	AutoCapture bool   `json:"autoCapture"`
	OrderInfo   string `json:"orderInfo"`
	RequestId   string `json:"requestId"`
	ExtraData   string `json:"extraData"`
	Signature   string `json:"signature"`
}

type PaymentResponse struct {
	PartnerCode  string `json:"partnerCode"`
	RequestId    string `json:"requestId"`
	OrderId      string `json:"orderId"`
	Amount       string `json:"amount"`
	ResponseTime string `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   string `json:"resultCode"`
	PayUrl       string `json:"payUrl"`
}

type PaymentHandler struct {
	orderService orderService.OrderService
}

func NewPaymentHandler(orderService *orderService.OrderService) *PaymentHandler {
	return &PaymentHandler{
		orderService: *orderService,
	}
}

func (h *PaymentHandler) Payment(c *gin.Context) {
	// Process body
	var shipping shippingModel.Shipping
	if err := c.ShouldBind(&shipping); c.Request.Body == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}

	// Get id
	id := c.Param("id")
	idVal, _ := strconv.Atoi(id)
	order, err := h.orderService.GetOrder(c, idVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "wrong id",
		})
		return
	}

	paygate := "https://test-payment.momo.vn"
	partnerCode := "MOMOBKUN20180529"
	accessKey := "klm05TvNBzhg7h7j"
	secretKey := "at67qH6mk8w5Y1nAyMoYKMWACiEi2bsa"

	cfg := config.GetConfig()

	// Get host of ngrok
	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "a",
		})
		return
	}
	defer resp.Body.Close()

	bodyRead, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "b",
		})
		return
	}
	bodyRes := gjson.Get(string(bodyRead), "tunnels.0.public_url")
	if !bodyRes.Exists() {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "c",
		})
		return
	}
	tunnelURL := bodyRes.String()
	notifyUrl := tunnelURL + "/checkout"

	returnUrl := cfg.HostFrontend

	lang := "vi"

	currTime := strconv.Itoa(int(time.Now().Unix()))
	requestId := currTime + "id"
	orderId := currTime + ":0123456778"
	autoCapture := true
	requestType := "captureWallet"
	amount := strconv.Itoa(order.TotalPrice)
	orderInfo := "Thanh toán qua ví MoMo"

	// Encoding body to extra data
	bodyStr, _ := json.Marshal(shipping)
	bodyEncB64 := b64.StdEncoding.EncodeToString(bodyStr)
	// extraData := "ew0KImVtYWlsIjogImh1b25neGRAZ21haWwuY29tIg0KfQ=="
	extraData := bodyEncB64
	signature := "accessKey=" + accessKey + "&amount=" + amount + "&extraData=" + extraData + "&ipnUrl=" + notifyUrl + "&orderId=" + orderId + "&orderInfo=" + orderInfo + "&partnerCode=" + partnerCode + "&redirectUrl=" + returnUrl + "&requestId=" + requestId + "&requestType=" + requestType

	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(signature))
	signature = hex.EncodeToString(hash.Sum(nil))

	payload := Payload{
		PartnerCode: partnerCode,
		PartnerName: "Test",
		StoreId:     partnerCode,
		RequestType: requestType,
		IpnUrl:      notifyUrl,
		RedirectUrl: returnUrl,
		OrderId:     orderId,
		Amount:      amount,
		Lang:        lang,
		AutoCapture: autoCapture,
		OrderInfo:   orderInfo,
		RequestId:   requestId,
		ExtraData:   extraData,
		Signature:   signature,
	}
	endPoint := "/v2/gateway/api/create"
	body, _ := json.Marshal(&payload)

	// Make request
	req, err := http.NewRequest("POST", paygate+endPoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Process response
	var result PaymentResponse
	r, _ := io.ReadAll(res.Body)

	_ = json.Unmarshal(r, &result)
	c.JSON(http.StatusOK, result)
}
