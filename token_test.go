package main

import (
	"testing"
)

func Test_GenGems(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "simple", input: "4e", want: "4e0s0r0d0o0j"},
		{name: "test2", input: "1e1s1o", want: "1e1s0r0d1o0j"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			want := tt.want
			gems := GenGems(tt.input)
			got := gems.String(true)
			if want != got {
				t.Errorf("test: %swanted: %s != got: %s", tt.name, want, got)
			}
		})
	}

}
