package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/influxdata/line-protocol/v2/lineprotocol"
)

func PostToPrometheus() error {
	userID := os.Getenv("PROM_USER_ID")
	apiKey := os.Getenv("PROM_API_KEY")
	baseUrl := os.Getenv("PROM_BASE_URL")

	if userID == "" || apiKey == "" || baseUrl == "" {
		return fmt.Errorf("missing environment variables, please set PROM_USER_ID, PROM_API_KEY, and PROM_BASE_URL")
	}

	fmt.Println("Prometheus (Influx Line Proto) Tester")
	url := baseUrl + "/write"

	body := assemble_sample_payload()

	// Prepare the request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
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
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Get the response status code
	statusCode := resp.StatusCode
	fmt.Printf("test writing data complete! returned status code %d\n", statusCode)

	return nil
}

// Assemble a line protocol payload
// * Line protocol https://docs.influxdata.com/influxdb/cloud/reference/syntax/line-protocol/#elements-of-line-protocol
// Prometheus-specific note: using `fieldKey` of `value` i.e, `value=123.45`, will have Grafana name the measurement as simply the Measurement name example: sandbox
// All other `fieldKey`s will be appended to the measurement name, separated by an underscore, example: `metric=123.45` -> `sandbox_metric` in Prometheus/Grafana
func assemble_sample_payload() []byte {
	var enc lineprotocol.Encoder

	// The name of the measurement
	enc.StartLine("sandbox")

	// Tags (labels)
	enc.AddTag("label1", "abc")
	// enc.AddTag("label2", "123")

	// Fields (key:value)
	// lineprotocol.MustNewValue() Acceptable data types:
	// * bool -> automatically converted to 0.0 or 1.0
	// * unsigned int
	// * signed int
	// * float
	enc.AddField("metric_1", lineprotocol.MustNewValue(true))

	// Timestamp
	// Note: Grafana only support ms precision - https://grafana.com/docs/grafana-cloud/data-configuration/metrics/metrics-influxdb/push-from-telegraf/#current-limitations
	enc.EndLine(time.Now())

	// Encode and check for errors
	if err := enc.Err(); err != nil {
		panic(fmt.Errorf("encoding error: %v", err))
	}

	fmt.Printf("encoded: %s", enc.Bytes())

	return enc.Bytes()
}
