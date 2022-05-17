package foo

import "testing"

func TestFoo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal",
			want: "Foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
