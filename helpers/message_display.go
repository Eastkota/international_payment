package helpers

func GetResponseDescription(responseCode string) string {
    switch responseCode {
    case "000":
        return "Success Transaction"
    case "001":
        return "Transaction failed"
    case "002":
        return "Transaction have been void/reversal"
    case "003":
        return "Transaction timeout"
    case "004":
        return "Transaction in processing"
    case "101":
        return "Message not recognised"
    case "102":
        return "Message Version Number received is not valid"
    case "201":
        return "A message element required as defined in Table A.1 is missing"
    case "203":
        return "Data element not in the required format or value is invalid"
    case "301":
        return "Transaction ID received is not valid"
    case "302":
        return "Data could not be decrypted due to technical reason"
    case "503":
        return "Invalid Merchant"
    case "5A0":
        return "MAC verification failed"
    case "999":
        return "Challenge Failed"
    default:
        return "Unknown Error (" + responseCode + ")"
    }
}