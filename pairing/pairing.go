package pairing

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
	
	"github.com/gertjaap/litbox-client/ssh"
)

type PairingRequest struct {
    PublicKey string
}

func Pair(torHost string, hmacKey string, torProxy string) {
	
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 30} 

	publicKey := ssh.GenerateKeyPair()

	request := PairingRequest{publicKey}

	jsonStr, err := json.Marshal(request)

	key := []byte(hmacKey)
	h := hmac.New(sha256.New, key)
	h.Write(jsonStr)
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	req, err := http.NewRequest("POST", torHost + "/pair", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-HMAC-Signature", signature);

    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error calling pair endpoint", err)
    }
	defer resp.Body.Close()
	
}