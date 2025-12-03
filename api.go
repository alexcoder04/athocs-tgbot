package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getLiveData() ([][]string, error) {
	resp, err := http.Get(Config.ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %w", err)
	}

	r := csv.NewReader(strings.NewReader(string(body)))
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("CSV parse error: %w", err)
	}

	return records, nil
}
