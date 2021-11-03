package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	org := "innotechdev"
	bucket := "innotechdev-bucket"
	token := "c747ea0f-b7ae-45f5-aa4c-14453065bb5f"

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:7076", token)

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)
	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())
	// write point immediately
	err := writeAPI.WritePoint(context.Background(), p)
	fmt.Println("Error:", err)

	// Get query client
	queryAPI := client.QueryAPI(org)
	// Get parser flux query result
	query := `from(bucket:"` + bucket + `")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		fmt.Println("Error:", err)
	}

	// Ensures background processes finishes
	client.Close()
}
