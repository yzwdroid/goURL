package main

import (
	"reflect"
	"testing"
)

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
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStatusFromLink(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStatusFromLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getStatusFromLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
