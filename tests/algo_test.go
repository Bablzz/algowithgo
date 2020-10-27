package algo_test

import (
	"github.com/bablzz/algowithgo/pkg/algo"
	fuzz "github.com/google/gofuzz"
	"reflect"
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
			name:    "Positive sum",
			args:    args{list: []int{1, 2, 3, 4, 5}},
			wantSum: 15,
		},
		{
			name:    "Negative sum",
			args:    args{list: []int{-1, -2, -3, -4, -5}},
			wantSum: -15,
		},
		{
			name:    "Zero sum",
			args:    args{list: []int{0, 0, 0, 0}},
			wantSum: 0,
		},
		{
			name:    "Empty sum",
			args:    args{list: []int{}},
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

func TestQuicksort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Two elements",
			args: args{arr: []int{5, 4}},
			want: []int{4, 5},
		},
		{
			name: "One elements",
			args: args{arr: []int{-10}},
			want: []int{-10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.Quicksort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuicksort2(t *testing.T) {
	type args struct {
		arr     []int
		less    int
		greater int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Two elements",
			args: args{arr: []int{5, 4}},
			want: []int{4, 5},
		},
		{
			name: "One elements",
			args: args{arr: []int{-10}},
			want: []int{-10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			algo.Quicksort2(tt.args.arr, 0, len(tt.args.arr)-1)
			result := reflect.DeepEqual(tt.args.arr, tt.want)
			if result {
				t.Logf("PASS, %v", tt.args.arr)
			} else {
				t.Errorf("Quicksort2() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Two elements",
			args: args{arr: []int{5, 4}},
			want: []int{4, 5},
		},
		{
			name: "One elements",
			args: args{arr: []int{-10}},
			want: []int{-10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr  []int
		elem int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Find element in left",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				elem: 4,
			},
			want: 3,
		},
		{
			name: "Find element in right",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				elem: 2,
			},
			want: 1,
		},
		{
			name: "Right boundary",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				elem: 1,
			},
			want: 0,
		},
		{
			name: "Left boundary",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				elem: 5,
			},
			want: 4,
		},
		{
			name: "Not found",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				elem: 8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.Bs(tt.args.arr, tt.args.elem); got != tt.want {
				t.Errorf("Bs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseNotation(t *testing.T) {
	type args struct {
		polish string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple",
			args: args{"2 + 2"},
			want: "2 2 +",
		},
		{
			name: "High proiritet first",
			args: args{"2 * 2 - 4"},
			want: "2 2 * 4 -",
		},
		{
			name: "High proiritet last",
			args: args{"2 + 2 * 3"},
			want: "2 2 3 * +",
		},
		{
			name: "Double equivalent",
			args: args{"2 + 2 - 2"},
			want: "2 2 + 2 -",
		},
		{
			name: "Without brackets",
			args: args{"2 * 2 - 3 + 2"},
			want: "2 2 * 3 - 2 +",
		},
		{
			name: "Simple with brackets",
			args: args{"1 * ( 2 - 3 )"},
			want: "1 2 3 - *",
		},
		{
			name: "Complicated with brackets",
			args: args{"1 + ( 2 - 3 ) * 4"},
			want: "1 2 3 - 4 * +",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.ReverseNotation(tt.args.polish); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse_notation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSets(t *testing.T) {
	sets := algo.Set{}
	sets.New()
	sets.Add("test")
	sets.Add("test")
	_, ok := sets.Get("test")
	if !ok {
		t.Errorf("Can't find element")
	}

	sets.Delete("test")
	_, ok = sets.Get("test")
	if ok {
		t.Errorf("Find exists after delete")
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GCD equal 6",
			args: args{
				num1: 54,
				num2: 24,
			},
			want: 6,
		},
		{
			name: "GCD equal 90",
			args: args{
				num1: 630,
				num2: 2700,
			},
			want: 90,
		},
		{
			name: "GCD equal 24",
			args: args{
				num1: 96,
				num2: 72,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.Gcd(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursionGcd(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GCD equal 6",
			args: args{
				num1: 54,
				num2: 24,
			},
			want: 6,
		},
		{
			name: "GCD equal 90",
			args: args{
				num1: 630,
				num2: 2700,
			},
			want: 90,
		},
		{
			name: "GCD equal 24",
			args: args{
				num1: 96,
				num2: 72,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.GCDRecur(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
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

func benchmarkQuicksort(l int, b *testing.B) {
	f := fuzz.New()
	var myInt []int
	var j int
	for i := 0; i < l; i++ {
		f.Fuzz(&j)
		myInt = append(myInt, j)
	}
	for n := 0; n < b.N; n++ {
		algo.Quicksort(myInt)
	}
}

func benchmarkQuicksort2(l int, b *testing.B) {
	f := fuzz.New()
	var myInt []int
	var j int
	for i := 0; i < l; i++ {
		f.Fuzz(&j)
		myInt = append(myInt, j)
	}
	for n := 0; n < b.N; n++ {
		algo.Quicksort2(myInt, 0, len(myInt)-1)
	}
}

func benchmarkMergeSort(l int, b *testing.B) {
	f := fuzz.New()
	var myInt []int
	var j int
	for i := 0; i < l; i++ {
		f.Fuzz(&j)
		myInt = append(myInt, j)
	}
	for n := 0; n < b.N; n++ {
		algo.MergeSort(myInt)
	}
}

func BenchmarkSumList1_10(b *testing.B)   { benchmarkSumList(10, b) }
func BenchmarkSumList2_20(b *testing.B)   { benchmarkSumList(20, b) }
func BenchmarkSumList3_100(b *testing.B)  { benchmarkSumList(100, b) }
func BenchmarkSumList4_1500(b *testing.B) { benchmarkSumList(1500, b) }

func BenchmarkReverseSumList1_10(b *testing.B)   { benchmarkReverseSumList(10, b) }
func BenchmarkReverseSumList1_20(b *testing.B)   { benchmarkReverseSumList(20, b) }
func BenchmarkReverseSumList1_100(b *testing.B)  { benchmarkReverseSumList(100, b) }
func BenchmarkReverseSumList1_1500(b *testing.B) { benchmarkReverseSumList(1500, b) }

func BenchmarkQuicksort1_10(b *testing.B)   { benchmarkQuicksort(10, b) }
func BenchmarkQuicksort1_20(b *testing.B)   { benchmarkQuicksort(20, b) }
func BenchmarkQuicksort1_100(b *testing.B)  { benchmarkQuicksort(100, b) }
func BenchmarkQuicksort1_1500(b *testing.B) { benchmarkQuicksort(1500, b) }

func BenchmarkQuicksort2_10(b *testing.B)   { benchmarkQuicksort(10, b) }
func BenchmarkQuicksort2_20(b *testing.B)   { benchmarkQuicksort(20, b) }
func BenchmarkQuicksort2_100(b *testing.B)  { benchmarkQuicksort(100, b) }
func BenchmarkQuicksort2_1500(b *testing.B) { benchmarkQuicksort(1500, b) }

func BenchmarkMergeSort1_10(b *testing.B)   { benchmarkMergeSort(10, b) }
func BenchmarkMergeSort1_20(b *testing.B)   { benchmarkMergeSort(20, b) }
func BenchmarkMergeSort1_100(b *testing.B)  { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1_1500(b *testing.B) { benchmarkMergeSort(1500, b) }
