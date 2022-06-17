package agent

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"

	"github.com/lainio/err2/try"
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
	const seedLength = 32
	seed := try.To1(generateRandomString(seedLength))
	go registerDIDToLedger(seed)
	return seed
}

func registerDIDToLedger(seed string) {
	payload := []byte(fmt.Sprintf(`{"seed":%q}`, seed))
	path := fmt.Sprintf("%s/register", os.Getenv("LEDGER_URL"))
	res := try.To1(doHTTPPostRequest(path, payload))
	var registerRes registerResponse
	if err := json.Unmarshal(res, &registerRes); err == nil {
		fmt.Println("Registered public DID", registerRes.Did, "and verkey", registerRes.Verkey)
		publicDID = registerRes.Did
		publicVerkey = registerRes.Verkey
	} else {
		fmt.Println("Error when registering DID:", string(res))
	}
}

func doHTTPPostRequest(url string, body []byte) ([]byte, error) {
	fmt.Println("post request", string(body), " to url:", url)
	req, err := http.NewRequestWithContext(context.TODO(), "POST", url, bytes.NewBuffer(body))
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

	return io.ReadAll(resp.Body)
}
