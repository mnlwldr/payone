# whats-this

A Simple PAYONE Integration in Golang 

## How to use it
```go
import "https://github.com/mnlwldr/payone"
```

## Usage
```go
response, err := sendRequest(parameters)
```

### Parameters

You can find the documentation for PAYONE [here](https://docs.payone.com/)

```go
parameters := url.Values{}

// classic payone accoint informations. 
parameters.Set("aid", "")
parameters.Set("mid", "")
parameters.Set("portalid", "")
parameters.Set("key", "") // the key has to be hashed as md5
parameters.Set("mode", "test") // can be "live" for actual transactions
parameters.Set("api_version", "3.8")
parameters.Set("encoding", "UTF-8")

// payment type parameters
parameters.Set("clearingtype", "wlt")
parameters.Set("wallettype", "PPE")
parameters.Set("request", "authorization")
parameters.Set("narrative_text", "Just a test")
parameters.Set("reference", fmt.Sprintf("%d", rand.Intn(100)))
parameters.Set("currency", "EUR")
parameters.Set("amount", "1234") // in cent
parameters.Set("successurl", "https://example.com/success")
parameters.Set("errorurl", "https://example.com/error")
parameters.Set("backurl", "https://example.com/back")

// personal data. Not everything is mandatory
parameters.Set("salutation", "Frau")
parameters.Set("title", "")
parameters.Set("firstname", "Maraike")
parameters.Set("lastname", "Musterfrau")
parameters.Set("street", "Musterstrasse")
parameters.Set("zip", "12345")
parameters.Set("city", "Somewhere over the rainbow")
parameters.Set("country", "DE")
parameters.Set("email", "maraike.musterfrau@example.com")
parameters.Set("language", "de")
parameters.Set("gender", "f")
```

## Response
```go
type Response struct {
	Status      string 
	TxId        string 
	UserId      string 
	RedirectUrl string 
	Token       string
}
```
[![Go Reference](https://pkg.go.dev/badge/github.com/mnlwldr/payone.svg)](https://pkg.go.dev/github.com/mnlwldr/payone)