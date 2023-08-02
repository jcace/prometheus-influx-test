package main

import (
	"context"
	"fmt"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influx"
)

const INFLUXDB_TOKEN string = ""

func InfluxPoc() {
	fmt.Println("Influx POC")

	// Create client
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influx.New(influx.Configs{
		HostURL:   url,
		AuthToken: INFLUXDB_TOKEN,
	})

	if err != nil {
		panic(err)
	}
	// Close client at the end and escalate error if present
	defer func(client *influx.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	database := "spade"
	data := map[string]map[string]interface{}{
		"point1": {
			"tag":    "metric1",
			"field1": "text",
			"field2": 23.16,
		},
		"point2": {
			"tag":    "metric2",
			"field1": "other_text",
			"field2": 30,
		},
	}

	// Write data
	for key := range data {
		point := influx.NewPointWithMeasurement("sandbox").
			AddTag("tag", data[key]["tag"].(string)).
			AddField(data[key]["field1"].(string), data[key]["field2"])

		if err := client.WritePoints(context.Background(), database, point); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second) // separate points by 1 second
	}

	fmt.Println("Done writing data!")

}
