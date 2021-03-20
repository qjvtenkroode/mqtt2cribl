package mqtt2cribl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func SendToCribl(msgReader io.Reader) ([]byte, error) {
	resp, err := http.Post(os.Getenv("MQTT2CRIBL_ENDPOINT"), "application/json", msgReader)
	if err != nil {
		fmt.Printf("Error - POST: %s\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error - Response Body: %s\n", err)
		return nil, err
	}

	return body, nil
}
