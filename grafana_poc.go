package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/influxdata/line-protocol/v2/lineprotocol"
)

// UPDATE IN CONNECTION INFO HERE
const ()

// ###############

func GrafanaPoc() {
	fmt.Println("Prometheus (Influx Line Proto) POC")
	url := baseUrl + "/write"

	// * Line protocol https://docs.influxdata.com/influxdb/cloud/reference/syntax/line-protocol/#elements-of-line-protocol
	// Prometheus-specific note: using `fieldKey` of `value` i.e, `value=123.45`, will have Grafana name the measurement as simply the Measurement name example: sandbox
	// All other `fieldKey`s will be appended to the measurement name, separated by an underscore, example: `metric=123.45` -> `sandbox_metric` in Prometheus/Grafana
	//// body := "sandbox,label1=abc,label2=123 metric_1=154.45,metric_2=32.42"

	body := assemble_line_protocol()

	// Prepare the request
	// req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
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
	fmt.Println("Done writing data!")
	fmt.Println(statusCode)
}

func assemble_line_protocol() []byte {
	var enc lineprotocol.Encoder
	// ! MUST set to Millisecond precision, Grafana cloud does not support any higher - https://grafana.com/docs/grafana-cloud/data-configuration/metrics/metrics-influxdb/push-from-telegraf/#current-limitations
	enc.SetPrecision(lineprotocol.Millisecond)
	enc.StartLine("sandbox")
	enc.AddTag("label1", "abc")
	enc.AddTag("label2", "123")
	enc.AddField("metric_1", lineprotocol.MustNewValue(123.45))
	enc.EndLine(time.Time{})
	if err := enc.Err(); err != nil {
		panic(fmt.Errorf("encoding error: %v", err))
	}

	fmt.Printf("encoded: %s", enc.Bytes())

	return enc.Bytes()
}
