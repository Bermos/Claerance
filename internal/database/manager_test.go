package database

import (
	"testing"
)

func TestSetDriver(t *testing.T) {
	type args struct {
		driver string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"OK", args{driver: "sqlite3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDriver(tt.args.driver)
		})
	}
}

func TestSetURI(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"OK", args{uri: "test.db"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SetURI(tt.args.uri); dbURI != tt.args.uri {
				t.Errorf("SetURI() = %v, want %v", dbURI, tt.args.uri)
			}
		})
	}
}
