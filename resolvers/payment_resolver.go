package resolvers

import (
	"payment_service/model"
	"payment_service/services"

	"encoding/json"
	
	"github.com/graphql-go/graphql"
)

type PaymentResolver struct {
	Services *services.PaymentService
}

func NewPaymentResolver(service *services.PaymentService) *PaymentResolver {
	return &PaymentResolver{Services: service}
}

// resolver.go
func (r *PaymentResolver) CreateInternationalPayment(p graphql.ResolveParams) (interface{}, error) {
    var paymentInput model.MercReqInput

    // 1. Map Arguments
    inputData := p.Args["input"].(map[string]interface{})
    bytes, _ := json.Marshal(inputData)
    json.Unmarshal(bytes, &paymentInput)

    // 2. Call Service
    result, err := r.Services.CreateInternationalPayment(p.Context, paymentInput)
    if err != nil {
        // Return the error inside the result structure
        return map[string]interface{}{
            "error": map[string]interface{}{
                "message": err.Error(),
                "code":    "SERVICE_ERROR",
            },
        }, nil
    }

    // 3. IMPORTANT: Convert Map to a Slice of Key-Value pairs
    // GraphQL cannot resolve a Go Map[string]string automatically
    var payloadList []map[string]string
    for k, v := range result.Payload {
        payloadList = append(payloadList, map[string]string{
            "key":   k,
            "value": v,
        })
    }

    // 4. Return the data at the TOP LEVEL (no wrapper)
    return map[string]interface{}{
        "acsURL":  result.AcsURL,
        "payload": payloadList,
        "error":   result.Error,
    }, nil
}