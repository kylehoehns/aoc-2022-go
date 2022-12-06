package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{name: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{name: "bvwbjplbgvbhsrlpgdmjqwftvncz"},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{name: "nppdvjthqldpwncqszvftbrmjlhg"},
			want: 6,
		},
		{
			name: "Example 4",
			args: args{name: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			want: 10,
		},
		{
			name: "Example 5",
			args: args{name: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.name); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{name: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			want: 19,
		},
		{
			name: "Example 2",
			args: args{name: "bvwbjplbgvbhsrlpgdmjqwftvncz"},
			want: 23,
		},
		{
			name: "Example 3",
			args: args{name: "nppdvjthqldpwncqszvftbrmjlhg"},
			want: 23,
		},
		{
			name: "Example 4",
			args: args{name: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			want: 29,
		},
		{
			name: "Example 5",
			args: args{name: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.name); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}

}
