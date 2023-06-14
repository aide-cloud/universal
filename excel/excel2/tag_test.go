package excel2

import (
	"reflect"
	"testing"
)

func Test_parseTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"test1", args{tag: "a:b"}, map[string]string{"a": "b"}, false},
		{"test2", args{tag: "a:b;c:d"}, map[string]string{"a": "b", "c": "d"}, false},
		{"test3", args{tag: "a:b;c"}, nil, true},
		{"test4", args{tag: "a:b;c:d:e"}, nil, true},
		{"test5", args{tag: "a:b;c:d;e:f"}, map[string]string{"a": "b", "c": "d", "e": "f"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTag(tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}
