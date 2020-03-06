package palindrom

import "testing"

func TestIs(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{text: ""},
			want: true,
		},
		{
			name: "one",
			args: args{text: "a"},
			want: true,
		},
		{
			name: "one unicode",
			args: args{text: "Ü"},
			want: true,
		},
		{
			name: "even",
			args: args{text: "abcddcba"},
			want: true,
		},
		{
			name: "uneven",
			args: args{text: "abcdcba"},
			want: true,
		},
		{
			name: "unicode",
			args: args{text: "abÜcddcÜba"},
			want: true,
		},
		{
			name: "negativ",
			args: args{text: "negativ"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.text); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
