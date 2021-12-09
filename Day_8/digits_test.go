package main

import (
	"advent_of_code/day8/input"
	"testing"
)

func TestCountSimpleNumbers(t *testing.T) {
	test_string := []byte("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\n" +
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\n" +
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\n" +
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\n" +
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\n" +
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\n" +
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\n" +
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\n" +
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\n" +
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |fgae cfgab fg bagce\n")

	type args struct {
		patterns []input.Pattern
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
		want3 int
	}{
		{
			name:  "Test Task1",
			args:  args{input.Parse(&test_string)},
			want:  8,
			want1: 6,
			want2: 5,
			want3: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := CountSimpleNumbers(tt.args.patterns)
			if got != tt.want {
				t.Errorf("CountSimpleNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CountSimpleNumbers() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CountSimpleNumbers() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("CountSimpleNumbers() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
