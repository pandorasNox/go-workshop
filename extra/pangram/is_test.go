package pangram

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
			want: false,
		},
		{
			name: "lower case alphabet",
			args: args{text: "abcdefghijklmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: "upper case alphabet",
			args: args{text: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			want: true,
		},
		{
			name: "invalid text",
			args: args{text: "invalid text"},
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
