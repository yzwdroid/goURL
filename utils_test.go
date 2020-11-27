package main

import (
	"reflect"
	"testing"
)

// func Test_checkStatus(t *testing.T) {
// 	type args struct {
// 		link     string
// 		failOnly bool
// 		wg       *sync.WaitGroup
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 		})
// 	}
// }

// func Test_checkStatusJSON(t *testing.T) {
// 	type args struct {
// 		link string
// 		ch   chan urlStatus
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 		})
// 	}
// }

// func Test_checkStatusNoColor(t *testing.T) {
// 	type args struct {
// 		link     string
// 		failOnly bool
// 		wg       *sync.WaitGroup
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 		})
// 	}
// }

// func Test_dataTelscope(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []byte
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := dataTelscope(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("dataTelscope() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
