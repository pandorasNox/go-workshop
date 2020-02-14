## solution

```
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}
```

```
func TestAverage(t *testing.T) {
	given := []float64{5, 5}
	expected := 5.0

	if average(given) != expected {
		t.Error("TEST_FAILED")
	}
}
```

```
func TestAverage(t *testing.T) {
	type args struct {
		list []float64
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "avergae of 5, 5",
			args: args{
				list: []float64{5, 5},
			},
			want: 5.0,
			// want:    6.0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := average(tt.args.list)
			// got, err := average(tt.args.list)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("average() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			if got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
```