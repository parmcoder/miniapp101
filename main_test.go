package main

import "testing"

func Test_match(t *testing.T) {
	type args struct {
		keyword        string
		excludeKeyword string
		text           string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Success1", args{keyword: "a", excludeKeyword: "", text: "abc"}, true},
		{"Fail1", args{keyword: "d", text: "abc"}, false},
		{"Fail2", args{keyword: "a.c", text: "abc"}, false},
		{"Success2", args{keyword: "a+c", text: "abc"}, true},
		{"Fail3", args{keyword: "a.c+a.b", text: "a.b.c"}, false},
		{"Fail4", args{keyword: "a+b", excludeKeyword: "c", text: "abc"}, false},
		{"Success3", args{keyword: "a+b", excludeKeyword: "c", text: "abd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.keyword, tt.args.excludeKeyword, tt.args.text); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
