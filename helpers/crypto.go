package helpers

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "fmt"
    "os"
    "strings"
)

func GenerateMac(message string, privateKeyPath string) (string, error) {
    keyData, err := os.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    block, _ := pem.Decode(keyData)
    if block == nil {
        return "", fmt.Errorf("failed to parse private key PEM")
    }
    parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if err != nil {
        // Fallback: if it's actually an old PKCS1 key, try that too
        parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
        if err != nil {
            return "", fmt.Errorf("failed to parse private key: %v", err)
        }
    }

    // Cast the parsed key to an RSA Private Key
    privKey, ok := parsedKey.(*rsa.PrivateKey)
    if !ok {
        return "", fmt.Errorf("not an RSA private key")
    }

    hashed := sha256.Sum256([]byte(message))
    signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
    if err != nil {
        return "", err
    }
    return base64.RawURLEncoding.EncodeToString(signature), nil
}

func FormatKeyToBase64Url(rawKey []byte) string {
    // 1. Decode the PEM block
    block, _ := pem.Decode(rawKey)
    var derBytes []byte
    if block != nil {
        derBytes = block.Bytes
    } else {
        // Fallback: If already stripped of PEM headers, try decoding standard B64
        s := strings.NewReplacer(
            "-----BEGIN PUBLIC KEY-----", "",
            "-----END PUBLIC KEY-----", "",
            "\n", "", "\r", "", " ", "",
        ).Replace(string(rawKey))
        
        var err error
        derBytes, err = base64.StdEncoding.DecodeString(s)
        if err != nil {
            return s // Return as is if all else fails
        }
    }
    
    // 2. Encode to Raw URL Base64 (No padding '=')
    return base64.RawURLEncoding.EncodeToString(derBytes)
}