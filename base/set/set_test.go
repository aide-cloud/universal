package set

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
	}{
		{
			name: "test",
			s:    New[int](),
			args: args{
				item: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.item)
			if !tt.s.Has(tt.args.item) {
				t.Errorf("Set.Add() = %v, want %v", tt.s, tt.args.item)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
	}{
		{
			name: "test",
			s:    New[int](),
			args: args{
				item: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.item)
			if tt.s.Has(tt.args.item) {
				t.Errorf("Set.Remove() = %v, want %v", tt.s, tt.args.item)
			}
		})
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want int
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
			want: 3,
		},
		{
			name: "test",
			s:    New[int](),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
		},
		{
			name: "test",
			s:    New[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if !tt.s.IsEmpty() {
				t.Errorf("Set.Clear() = %v, want %v", tt.s, tt.s.IsEmpty())
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want []int
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
			want: []int{1, 2, 3},
		},
		{
			name: "test",
			s:    New[int](),
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ForEach(t *testing.T) {
	type args struct {
		f func(int)
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
			args: args{
				f: func(i int) {
					if i < 0 && i > 3 {
						t.Errorf("Set.ForEach() = %v, want %v", i, 1)
					}
				},
			},
		},
		{
			name: "test",
			s:    New[int](),
			args: args{
				f: func(i int) {
					t.Errorf("Set.ForEach() = %v, want %v", i, 1)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.ForEach(tt.args.f)
		})
	}
}

func TestSet_Map(t *testing.T) {
	type args struct {
		f func(int) int
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
		want Set[int]
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
			args: args{
				f: func(i int) int {
					return i + 1
				},
			},
			want: New[int](2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Filter(t *testing.T) {
	type args struct {
		f func(int) bool
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
		want Set[int]
	}{
		{
			name: "test",
			s:    New[int](1, 2, 3),
			args: args{
				f: func(i int) bool {
					return i > 1
				},
			},
			want: New[int](2, 3),
		},
		{
			name: "test",
			s:    New[int](),
			args: args{
				f: func(i int) bool {
					return i > 1
				},
			},
			want: New[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Filter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
