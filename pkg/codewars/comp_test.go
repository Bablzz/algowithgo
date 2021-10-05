package codewars

import "testing"

func TestComp(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive case",
			args: args{
				arr1: []int{121, 144, 19, 161, 19, 144, 19, 11},
				arr2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: true,
		},
		{
			name: "Negative digit",
			args: args{
				arr1: []int{-39, -13, 18, 82, 76, -45, 66, -15, 76, -99, -3, -43, 66, 2},
				arr2: []int{4, 9, 169, 225, 324, 1521, 1849, 2025, 4356, 4356, 5776, 5776, 6724, 9801},
			},
			want: true,
		},
		{
			name: "Empty arrays",
			args: args{
				arr1: []int{},
				arr2: []int{},
			},
			want: true,
		},
		{
			name: "Nil arrays",
			args: args{
				arr1: nil,
				arr2: nil,
			},
			want: false,
		},
		{
			name: "Kata arrays",
			args: args{
				arr1: []int{11, 19, 19, 19, 121, 144, 144, 161},
				arr2: []int{121, 361, 361, 14641, 20736, 20736, 25921, 36100},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comp(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("Comp() = %v, want %v", got, tt.want)
			}
		})
	}
}
