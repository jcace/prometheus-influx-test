package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

// UPDATE IN CONNECTION INFO HERE
const (
	userID  = ""
	apiKey  = ""
	baseUrl = ""
)

// ###############

func main() {
	// * Line protocol https://docs.influxdata.com/influxdb/cloud/reference/syntax/line-protocol/#elements-of-line-protocol
	// Prometheus-specific note: using `fieldKey` of `value` i.e, `value=123.45`, will have Grafana name the measurement as simply the Measurement name example: sandbox
	// All other `fieldKey`s will be appended to the measurement name, separated by an underscore, example: `metric=123.45` -> `sandbox_metric` in Prometheus/Grafana
	body := "sandbox,label1=abc,label2=123 metric_1=154.45,metric_2=32.42"

	// Prepare the request
	url := baseUrl + "/write"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "text/plain")

	// Encode the authentication credentials (Basic Auth)
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userID, apiKey)))
	req.Header.Set("Authorization", "Basic "+auth)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Get the response status code
	statusCode := resp.StatusCode
	fmt.Println(statusCode)
}
