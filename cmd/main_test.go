package main

import "testing"

func TestExampleScrape(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"ok",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExampleScrape(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleScrape() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
