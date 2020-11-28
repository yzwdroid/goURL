package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExtractURL(t *testing.T) {
	result := extractURL("https://www.google.ca, https://www.facebook.com")

	expected := []string{"https://www.google.ca", "https://www.facebook.com"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}

func Test_removeDuplicate(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"name", args{[]string{"http://www.google.ca", "http://www.google.ca"}}, []string{"http://www.google.ca"}},
		{"name", args{[]string{"http://www.google.ca", "http://zyang.ca", "http://www.google.ca"}}, []string{"http://www.google.ca", "http://zyang.ca"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicate(tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStatusFromLink(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	}))
	defer ts.Close()
	url := ts.URL

	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
		{"ok", args{url}, 200, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStatusFromLink(tt.args.link)
			if err != nil {
				t.Errorf("getStatusFromLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getStatusFromLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_404_getStatusFromLink(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}))
	defer ts.Close()
	url := ts.URL

	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
		{"ok", args{url}, 404, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStatusFromLink(tt.args.link)
			if err != nil {
				t.Errorf("getStatusFromLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getStatusFromLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_300_getStatusFromLink(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(300)
		w.Write([]byte("Not Found"))
	}))
	defer ts.Close()
	url := ts.URL

	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
		{"ok", args{url}, 300, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStatusFromLink(tt.args.link)
			if err != nil {
				t.Errorf("getStatusFromLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getStatusFromLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
