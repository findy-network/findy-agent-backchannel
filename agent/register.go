package agent

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	"github.com/lainio/err2"
)

// TODO: thread-safe
var (
	publicDID    = "Y9oNbFNTgxrRuvxQk3xEzr"
	publicVerkey = "HymVhRozF7o9Sh9iyKXu5WKHP95YERhrpZxGx5d6WhYw"
)

type registerResponse struct {
	Did    string `json:"did"`
	Seed   string `json:"seed"`
	Verkey string `json:"verkey"`
}

func PublicDID() string {
	return publicDID
}

func PublicVerkey() string {
	return publicVerkey
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

func registerDID() string {
	seed := err2.String.Try(generateRandomString(32))
	go registerDIDToLedger(seed)
	return seed
}

func registerDIDToLedger(seed string) {
	payload := []byte(fmt.Sprintf(`{"seed":"%s"}`, seed))
	path := fmt.Sprintf("%s/register", os.Getenv("LEDGER_URL"))
	res := err2.Bytes.Try(doHttpPostRequest(path, payload))
	var registerResponse registerResponse
	if err := json.Unmarshal(res, &registerResponse); err == nil {
		publicDID = registerResponse.Did
		publicVerkey = registerResponse.Verkey
	} else {
		fmt.Println("Error when registering DID:", string(res))
	}
}

func doHttpPostRequest(url string, body []byte) ([]byte, error) {
	fmt.Println("post request", string(body), " to url:", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
