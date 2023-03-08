package binary

import (
	"testing"
)

func TestBinarySearchWithSort(t *testing.T) {
	type args struct {
		list []string
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true test",
			args: args{
				list: []string{"a", "b", "c"},
				word: "b",
			},
			want: true,
		},
		{
			name: "false test",
			args: args{
				list: []string{"a", "b", "c"},
				word: "d",
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchWithSort(tt.args.list, tt.args.word); got != tt.want {
				t.Errorf("BinarySearchWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
