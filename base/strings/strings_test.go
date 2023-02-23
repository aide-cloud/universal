package strings

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "non-empty string",
			args: args{s: "foo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmptyTrimmed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "non-empty string",
			args: args{s: "foo"},
			want: false,
		},
		{
			name: "string with whitespace",
			args: args{s: " foo "},
			want: false,
		},
		{
			name: "string with only whitespace",
			args: args{s: " "},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyTrimmed(tt.args.s); got != tt.want {
				t.Errorf("IsEmptyTrimmed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUTF8Len(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "string with 1 rune",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "string with 2 runes",
			args: args{s: "ab"},
			want: 2,
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好"},
			want: 2,
		},
		{
			name: "string with 2 chinese and 1 english",
			args: args{s: "你好a"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UTF8Len(tt.args.s); got != tt.want {
				t.Errorf("UTF8Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUTF8LenTrimmed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "string with 1 rune",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "string with 2 runes",
			args: args{s: "ab"},
			want: 2,
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好"},
			want: 2,
		},
		{
			name: "string with 2 chinese and 1 english",
			args: args{s: "你好a"},
			want: 3,
		},
		{
			name: "string with whitespace",
			args: args{s: " foo "},
			want: 3,
		},
		{
			name: "string with only whitespace",
			args: args{s: " "},
			want: 0,
		},
		{
			name: "string with only whitespace and chinese",
			args: args{s: " 你 好 "},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UTF8LenTrimmed(tt.args.s); got != tt.want {
				t.Errorf("UTF8LenTrimmed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLenTrimmed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "string with 1 rune",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "string with 2 runes",
			args: args{s: "ab"},
			want: 2,
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好"},
			want: 6,
		},
		{
			name: "string with 2 chinese and 1 english",
			args: args{s: "你好a"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenTrimmed(tt.args.s); got != tt.want {
				t.Errorf("LenTrimmed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string with 1 rune",
			args: args{s: "A"},
			want: "a",
		},
		{
			name: "string with 2 runes",
			args: args{s: "Ab"},
			want: "ab",
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好"},
			want: "你好",
		},
		{
			name: "string with 2 chinese and 1 english",
			args: args{s: "你好A"},
			want: "你好a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.s); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string with 1 rune",
			args: args{s: "a"},
			want: "A",
		},
		{
			name: "string with 2 runes",
			args: args{s: "ab"},
			want: "AB",
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好AbCd"},
			want: "你好ABCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTitle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string with 1 rune",
			args: args{s: "a"},
			want: "A",
		},
		{
			name: "string with 2 runes",
			args: args{s: "ab"},
			want: "AB",
		},
		{
			name: "string with 2 chinese",
			args: args{s: "你好A_b_C_d"},
			want: "你好A_B_C_D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTitle(tt.args.s); got != tt.want {
				t.Errorf("ToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
