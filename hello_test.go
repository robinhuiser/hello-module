package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"t1", "Hello, world."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeOneSays(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{name: "Robin"}, "Robin says: Hello, world."},
		{"t2", args{name: "Maria"}, "Maria says: Hello, world."},
		{"t3", args{name: "David"}, "David says: Hello, world."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeOneSays(tt.args.name); got != tt.want {
				t.Errorf("SomeOneSays() = %v, want %v", got, tt.want)
			}
		})
	}
}
