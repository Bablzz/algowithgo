package algo_test

import (
	"github.com/bablzz/algowithgo/pkg/algo"
	fuzz "github.com/google/gofuzz"
	"testing"
)

func Test_NumInList(t *testing.T) {
	type args struct {
		list []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has in list",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				num:  3,
			},
			want: true,
		},
		{
			name: "has in list border",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				num:  1,
			},
			want: true,
		},
		{
			name: "has in list border",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				num:  5,
			},
			want: true,
		},
		{
			name: "has no in list",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				num:  12,
			},
			want: false,
		},
		{
			name: "empty list",
			args: args{
				list: []int{},
				num:  0,
			},
			want: false,
		},
		{
			name: "zero in list",
			args: args{
				list: []int{0},
				num:  0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.NumInList(tt.args.list, tt.args.num); got != tt.want {
				t.Errorf("NumInList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Sum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "Positive sum",
			args:args{list:[]int{1,2,3,4,5}},
			wantSum: 15,
		},
		{
			name: "Negative sum",
			args:args{list:[]int{-1,-2,-3,-4,-5}},
			wantSum: -15,
		},
		{
			name: "Zero sum",
			args:args{list:[]int{0, 0, 0, 0}},
			wantSum: 0,
		},
		{
			name: "Empty sum",
			args:args{list:[]int{}},
			wantSum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := algo.SumList(tt.args.list); gotSum != tt.wantSum {
				t.Errorf("sum_list() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := algo.ReverseSum(tt.args.list); gotSum != tt.wantSum {
				t.Errorf("sum_list() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}


func benchmarkSumList(l int, b *testing.B) {
	f := fuzz.New()
	var myInt []int
	var j int
	for i := 0; i < l; i++ {
		f.Fuzz(&j)
		myInt = append(myInt, j)
	}
	for n := 0; n < b.N; n++ {
		algo.SumList(myInt)
	}
}

func benchmarkReverseSumList(l int, b *testing.B) {
	f := fuzz.New()
	var myInt []int
	var j int
	for i := 0; i < l; i++ {
		f.Fuzz(&j)
		myInt = append(myInt, j)
	}
	for n := 0; n < b.N; n++ {
		algo.ReverseSum(myInt)
	}
}

func BenchmarkSumList1_10(b *testing.B)  { benchmarkSumList(10, b) }
func BenchmarkSumList2_20(b *testing.B)  { benchmarkSumList(20, b) }
func BenchmarkSumList3_100(b *testing.B)  { benchmarkSumList(20, b) }
func BenchmarkSumList4_1500(b *testing.B)  { benchmarkSumList(20, b) }

func BenchmarkReverseSumList1_10(b *testing.B)  { benchmarkReverseSumList(10, b) }
func BenchmarkReverseSumList1_20(b *testing.B)  { benchmarkReverseSumList(20, b) }
func BenchmarkReverseSumList1_100(b *testing.B)  { benchmarkReverseSumList(20, b) }
func BenchmarkReverseSumList1_1500(b *testing.B)  { benchmarkReverseSumList(20, b) }