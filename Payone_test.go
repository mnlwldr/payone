package payone

import (
	"fmt"
	"math/rand"
	"net/url"
	"testing"
)

func TestSendRequest(t *testing.T) {

	parameters := url.Values{}
	// defautl parameters
	parameters.Set("aid", "")
	parameters.Set("mid", "")
	parameters.Set("portalid", "")
	parameters.Set("key", "")
	parameters.Set("mode", "")
	parameters.Set("api_version", "")
	parameters.Set("encoding", "")

	// parameters (PayPal)
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

	// Personal data
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

	got, err := sendRequest(parameters)
	if err != nil {
		t.Error(err)
	}
	if got.Status != "REDIRECT" {
		t.Errorf("TestSendRequest() = %q, want %q", got.Status, "REDIRECT")
	}
}
