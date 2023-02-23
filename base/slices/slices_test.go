package slices

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "mid",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 3,
			},
			want: 2,
		},
		{
			name: "not found",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 6,
			},
			want: -1,
		},
		{
			name: "first",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 1,
			},
			want: 0,
		},
		{
			name: "last",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "mid",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
		{
			name: "first",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 1,
			},
			want: true,
		},
		{
			name: "last",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	type args struct {
		a         []Student
		predicate func(Student) bool
	}
	tests := []struct {
		name string
		args args
		want []Student
	}{
		{
			name: "filter",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				predicate: func(s Student) bool {
					return s.Age > 18
				},
			},
			want: []Student{
				{Name: "小花", Age: 19},
				{Name: "小北", Age: 20},
			},
		},
		{
			name: "filter all",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				predicate: func(s Student) bool {
					return s.Age > 0
				},
			},
			want: []Student{
				{Name: "小明", Age: 18},
				{Name: "小花", Age: 19},
				{Name: "小北", Age: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.a, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		a []int
		f func(int) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "map",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x int) float64 {
					return float64(x) * 2
				},
			},
			want: []float64{2, 4, 6, 8, 10},
		},
		{
			name: "map all",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x int) float64 {
					return float64(x)
				},
			},
			want: []float64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTo(t *testing.T) {
	type args struct {
		a []int
		f func(int) int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "to",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				f: func(x int) int64 {
					return int64(x)
				},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("To() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		a    []int
		less func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort",
			args: args{
				a: []int{5, 4, 3, 2, 1},
				less: func(i, j int) bool {
					return i > j
				},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.a, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortStruct(t *testing.T) {
	type Student struct {
		Name    string
		Age     int
		Math    float32
		Chinese float32
		English float32
	}
	type args struct {
		a    []Student
		less func(i, j Student) bool
	}
	tests := []struct {
		name string
		args args
		want []Student
	}{
		{
			name: "sort age",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
					{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
					{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				},
				less: func(i, j Student) bool {
					return i.Age > j.Age
				},
			},
			want: []Student{
				{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
				{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
			},
		},
		{
			name: "sort math",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
					{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
					{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				},
				less: func(i, j Student) bool {
					return i.Math > j.Math
				},
			},
			want: []Student{
				{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
				{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
				{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
			},
		},
		{
			name: "sort chinese",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
					{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
					{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				},
				less: func(i, j Student) bool {
					return i.Chinese < j.Chinese
				},
			},
			want: []Student{
				{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
				{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
			},
		},
		{
			name: "sort english",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
					{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
					{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				},
				less: func(i, j Student) bool {
					return i.English < j.English
				},
			},
			want: []Student{
				{Name: "小北", Age: 20, Math: 70, Chinese: 60, English: 50},
				{Name: "小花", Age: 19, Math: 80, Chinese: 70, English: 60},
				{Name: "小明", Age: 18, Math: 90, Chinese: 80, English: 70},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.a, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverse",
			args: args{
				a: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "reverse",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "unique",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "unique",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

type Student struct {
	Name string
	Age  int
}

func (l Student) UniqueKey() string {
	return l.Name
}

func TestUniqueStruct(t *testing.T) {
	type args struct {
		a []Student
	}
	tests := []struct {
		name string
		args args
		want []Student
	}{
		{
			name: "unique",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
			},
			want: []Student{
				{Name: "小明", Age: 18},
				{Name: "小花", Age: 19},
				{Name: "小北", Age: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStruct(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexStruct(t *testing.T) {
	type args struct {
		a []Student
		f func(Student) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "index",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Name == "小北"
				},
			},
			want: 2,
		},
		{
			name: "index",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Age == 19
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexStruct(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("IndexStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStruct(t *testing.T) {
	type args struct {
		a []Student
		f func(Student) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Name == "小北"
				},
			},
			want: true,
		},
		{
			name: "contains",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Age == 19
				},
			},
			want: true,
		},
		{
			name: "contains",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Age == 21
				},
			},
			want: false,
		},
		{
			name: "contains",
			args: args{
				a: []Student{
					{Name: "小明", Age: 18},
					{Name: "小花", Age: 19},
					{Name: "小北", Age: 20},
				},
				f: func(s Student) bool {
					return s.Name == "小北北"
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsStruct(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("ContainsStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUnique(t *testing.T) {
	type args struct {
		a [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge",
			args: args{
				a: [][]int{
					{1, 2, 3},
					{2, 3, 4},
					{3, 4, 5},
				},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUnique(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		a [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge",
			args: args{
				a: [][]int{
					{1, 2, 3},
					{2, 3, 4},
					{3, 4, 5},
				},
			},
			want: []int{1, 2, 3, 2, 3, 4, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUniqueStruct(t *testing.T) {
	type args struct {
		a [][]Student
	}
	tests := []struct {
		name string
		args args
		want []Student
	}{
		{
			args: args{
				a: [][]Student{
					{
						{Name: "小明", Age: 18},
						{Name: "小花", Age: 19},
						{Name: "小北", Age: 20},
						{Name: "小北", Age: 22},
						{Name: "小北", Age: 21},
					},
				},
			},
			want: []Student{
				{Name: "小明", Age: 18},
				{Name: "小花", Age: 19},
				{Name: "小北", Age: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUniqueStruct(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUniqueStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
