package config

import (
	"os"
)

func PostgresUser() string     { return os.Getenv("POSTGRES_USER") }
func PostgresPassword() string { return os.Getenv("POSTGRES_PASSWORD") }
func PostgresHost() string     { return os.Getenv("POSTGRES_HOST") }
func PostgresPort() string     { return os.Getenv("POSTGRES_PORT") }
func PostgresDB() string       { return os.Getenv("POSTGRES_DB") }

const (
    // Database                  = "educareRmaPayment"
	// PaymentResponseCollection = "payment_responses"
	// BENF_ID                   = "BE10000169"
	// BENF_BANK_CODE            = "01"
	// CURRENCY                  = "BTN"
	PAYMENT_PRIVATE_KEY       = "keys/private_key.pem"
	PAYMENT_PUBLIC_KEY        = "keys/public_key.pem"
	MK_REQUEST_URL            = "https://3dsecure.bob.bt/3dss/mkReq"
	MERC_REQ_URL              = "https://3dsecure.bob.bt/3dss/mercReq"
	// RMA_PAYMENT_API           = "https://bfssecure.rma.org.bt/BFSSecure/nvpapi"
	Database                = "educare3DS"
    PaymentIntentCollection = "payment_intents"
)

func MembershipApi() string  { return os.Getenv("MEMBERSHIP_API") }
func AuthServiceApi() string { return os.Getenv("AUTH_SERVICE_API") }

// 3DS Merchant configuration functions
func MerchantId() string     { return os.Getenv("MERCHANT_ID") }
func MerchantSecret() string { return os.Getenv("MERCHANT_SECRET") }
func MerchantApiUrl() string { return os.Getenv("MERCHANT_API_URL") }
func MkRequestUrl() string   { return os.Getenv("MK_REQUEST_URL") }
func MercReqUrl() string     { return os.Getenv("MERC_REQ_URL") }

