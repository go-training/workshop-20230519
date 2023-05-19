package main

import (
	"testing"
)

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test01",
			want: "歡迎來到 Go 世界",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHello(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hello()
	}
}

func Test_foobar(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing for foobar name",
			args: args{
				name: "foobar",
			},
			want: "Welcome to foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foobar(tt.args.name); got != tt.want {
				t.Errorf("foobar() = %v, want %v", got, tt.want)
			}
		})
	}
}
