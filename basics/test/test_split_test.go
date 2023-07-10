package test

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		// TODO: Add test cases.
		{
			name: "字符串切割测试 1",
			args: struct {
				s   string
				sep string
			}{s: "this is bad good bad.", sep: "bad"},
			wantResult: []string{"this is ", " good ", "."},
		},
		{
			name: "字符串切割测试 2",
			args: args{
				s:   "abbcdibbdis",
				sep: "bb",
			},
			wantResult: []string{"a", "cdi", "dis"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 因为 slice 不能比较直接，借助反射包中的 DeepEqual 方法比较
			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %#v, want %#v", gotResult, tt.wantResult)
			}
		})
	}
}
