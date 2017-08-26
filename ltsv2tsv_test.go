package ltsv2tsv

import (
	"bytes"
	"reflect"
	"testing"
)

var converterTests = []struct {
	in  string
	out interface{}
}{
	{
		`
test1:hoge	test2:hoge2	test3:hoge3
test1:hoge4	test2:hoge5	test3:hoge6
test2:hoge7	test3:hoge8	test4:hoge9
test1:08:01:01
		`,
		[][]string{
			{"test1", "test2", "test3", "test4"},
			{"hoge", "hoge2", "hoge3", ""},
			{"hoge4", "hoge5", "hoge6", ""},
			{"", "hoge7", "hoge8", "hoge9"},
			{"08:01:01","","",""},
		},
	},
	{
		`
test1:hoge	test2:	test3:hoge3
test1:hoge4	test2:hoge5	test3:
test2:hoge7	test3:hoge8	test4:hoge9
		`,
		[][]string{
			{"test1", "test2", "test3", "test4"},
			{"hoge", "", "hoge3", ""},
			{"hoge4", "hoge5", "", ""},
			{"", "hoge7", "hoge8", "hoge9"},
		},
	},
	{
		`
test1:hoge	:hoge2	test3:hoge3
		`,
		[][]string{
			{"test1", "", "test3"},
			{"hoge", "hoge2", "hoge3"},
		},
	},
}

func TestConverter(t *testing.T) {
	for _, v := range converterTests {
		reader := NewConverter(bytes.NewBufferString(v.in))
		out, e := reader.Converter()
		if e != nil {
			t.Errorf("expected Read got %v", e)
			continue
		}
		if reflect.DeepEqual(out, v.out) == false {
			t.Errorf("aaa")
			t.Errorf("Convert(%q) \n %d \n want:\n %d", v.in, out, v.out)
			continue
		}
	}
}
