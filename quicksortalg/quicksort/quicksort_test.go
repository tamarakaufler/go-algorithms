package quicksort

import (
	"reflect"
	"testing"
)

func TestSortMeQuick(t *testing.T) {
	type args struct {
		si []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Sort array of 4 elems",
			args: args{
				si: []int{4, 2, 3, 5},
			},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Sort array of 5 elems",
			args: args{
				si: []int{2, 5, 3, 9, 7},
			},
			want: []int{2, 3, 5, 7, 9},
		},
		{
			name: "Sort array of 7 elems",
			args: args{
				si: []int{4, 2, 3, 5, 1, 9, 7},
			},
			want: []int{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name: "Sort array of duplicate elems",
			args: args{
				si: []int{4, 2, 3, 5, 1, 2, 7},
			},
			want: []int{1, 2, 2, 3, 4, 5, 7},
		},
		{
			name: "Sort array of duplicate elems (3x)",
			args: args{
				si: []int{4, 2, 3, 5, 2, 1, 2, 7},
			},
			want: []int{1, 2, 2, 2, 3, 4, 5, 7},
		},
		{
			name: "Sort array of duplicate elems - mixture",
			args: args{
				si: []int{4, 2, 3, 5, 2, 1, 2, 7, 1},
			},
			want: []int{1, 1, 2, 2, 2, 3, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortMeQuick(tt.args.si); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMeQuick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty array",
			args: args{
				l: []int{},
			},
			want: true,
		},
		{
			name: "one element array",
			args: args{
				l: []int{1},
			},
			want: true,
		},
		{
			name: "all elements the same",
			args: args{
				l: []int{1, 1, 1, 1, 1},
			},
			want: true,
		},
		{
			name: "first element different",
			args: args{
				l: []int{2, 1, 1, 1, 1},
			},
			want: false,
		},
		{
			name: "internal element different",
			args: args{
				l: []int{1, 1, 3, 1, 1},
			},
			want: false,
		},
		{
			name: "last element different",
			args: args{
				l: []int{1, 1, 1, 1, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.l); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
