package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"payment_service/config"
	"payment_service/helpers"
	"payment_service/model"
	"payment_service/repositories"
)

type PaymentService struct {
	Repository repositories.Repository
}

func NewPaymentService(repository repositories.Repository) *PaymentService {
	return &PaymentService{Repository: repository}
}

func SendMKRequest(ctx context.Context) (*model.MPIKeyRequest, error) {
    // 1. Read your public key
    pubKeyBytes, err := os.ReadFile(config.PAYMENT_PUBLIC_KEY)
    if err != nil {
        return nil, helpers.WrapInternal("processing payment", err)
    }

    // Bank requires RSA 2048 encoded in Base64Url
    formattedPubKey := helpers.FormatKeyToBase64Url(pubKeyBytes)

    merchantID := config.MerchantId()
    fmt.Println("Using Merchant ID:", merchantID)
    // Transaction ID must be minimum 6 digits
    purchaseID := fmt.Sprintf("REQ%d", time.Now().Unix())

    // 2. Generate MAC
    // Data string: merchantId + purchaseId + pubKey
    // dataToSign := merchantID + purchaseID + formattedPubKey

    // 3. Build Request matching the Sample Request in your document
    reqBody := model.MPIKeyRequest{
        MerchantID: merchantID,
        PurchaseID: purchaseID,
        PubKey:     formattedPubKey,
    }

    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        return nil, helpers.WrapInternal("processing payment", err)
    }

    // Log this to confirm it looks exactly like the bank's sample
    fmt.Println("SENDING JSON:", string(jsonData))

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Post(config.MK_REQUEST_URL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, helpers.WrapInternal("processing payment", err)
    }
    defer resp.Body.Close()

    fmt.Println("RECEIVED RESPONSE STATUS:", resp.Status)

    // 4. Decode Response
    var rawResponse map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
        return nil, helpers.WrapInternal("reading bank response", err)
    }

    // Check errorCode from your list (e.g., 201, 5A0)
    if code, ok := rawResponse["errorCode"].(string); ok && code != "000" {
        description := helpers.GetResponseDescription(code)
        return nil, helpers.NewInternalError(fmt.Sprintf("Bank Error %s: %s", code, description), nil)
    }

    // Map successful response (echoes merchantId/purchaseId and returns bank pubKey)
    var mkResponse model.MPIKeyRequest
    b, err := json.Marshal(rawResponse)
    if err != nil {
        return nil, helpers.WrapInternal("processing bank response", err)
    }
    if err := json.Unmarshal(b, &mkResponse); err != nil {
        return nil, helpers.WrapInternal("processing bank response", err)
    }

    return &mkResponse, nil
}

// CreateInternationalPayment handles the standard payment flow

func (s *PaymentService) CreateInternationalPayment(ctx context.Context, inputData model.MercReqInput) (*model.InternationalPaymentResult, error) {
    txID := fmt.Sprintf("%d", time.Now().Unix())
    purchaseDate := time.Now().Format("20060102150405")
    currencyCode := "840"

    // 1. Generate MAC
    dataToSign := config.MerchantId() + txID + inputData.TransactionAmount + currencyCode + purchaseDate + "SALES"
    mac, err := helpers.GenerateMac(dataToSign, config.PAYMENT_PRIVATE_KEY)
    if err != nil {
        return nil, helpers.WrapInternal("processing payment", err)
    }

    // 2. Fix the "not enough arguments" error here
    err = s.Repository.CreateMercRequest(ctx, inputData, txID, purchaseDate, mac)
    if err != nil {
        return nil, helpers.WrapInternal("saving payment details", err)
    }

    // 3. This matches the struct we fixed in Step 1
    return &model.InternationalPaymentResult{
        AcsURL: config.MERC_REQ_URL,
        Payload: map[string]string{
            "MPI_MERC_ID":    config.MerchantId(),
            "MPI_TRXN_ID":    txID,
            "MPI_PURCH_AMT":  inputData.TransactionAmount,
            "MPI_PURCH_CURR": currencyCode,
            "MPI_PURCH_DATE": purchaseDate,
            "MPI_TRANS_TYPE": "SALES",
            "MPI_MAC":        mac,
        },
    }, nil
}
