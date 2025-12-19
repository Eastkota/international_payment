package helpers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

func ConvertInsertedIDToString(insertedID interface{}) string {
	switch v := insertedID.(type) {
	case primitive.ObjectID: // If it's a MongoDB ObjectID
		return v.Hex()
	case string: // If it's already a string
		return v
	default: // Fallback for unexpected types
		return fmt.Sprintf("%v", v)
	}
}
