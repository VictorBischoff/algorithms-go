package linear

import "testing"

func TestSimpleLinearSearch(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true test",
			args: args{
				list: []int{1,2,3,4,5,6,7,8,9,10},
				n: 5,
			},
			want: true,
		},
		{
			name: "false test",
			args: args{
				list: []int{1,2,3,4,5,6,7,8,9,10},
				n: 11,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleLinearSearch(tt.args.list, tt.args.n); got != tt.want {
				t.Errorf("SimpleLinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
