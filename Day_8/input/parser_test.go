package input

import (
	"advent_of_code/day8/signal"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	test_string := []byte("acedgfb cdfbe ab | cdfeb fcadb\n")
	type args struct {
		input *[]byte
	}
	tests := []struct {
		name string
		args args
		want []signal.Pattern
	}{
		{
			name: "Parse Test",
			args: args{&test_string},
			want: []signal.Pattern{
				signal.Pattern{
					S: []signal.Signal{
						signal.Signal{V: signal.A | signal.C | signal.E | signal.D | signal.G | signal.F | signal.B},
						signal.Signal{V: signal.C | signal.D | signal.F | signal.B | signal.E},
						signal.Signal{V: signal.A | signal.B},
					},
					O: []signal.Signal{
						signal.Signal{V: signal.C | signal.D | signal.F | signal.E | signal.B},
						signal.Signal{V: signal.F | signal.C | signal.A | signal.D | signal.B},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
