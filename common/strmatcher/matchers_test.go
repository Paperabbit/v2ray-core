package strmatcher_test

import (
	"testing"

	"v2ray.com/core/common"
	. "v2ray.com/core/common/strmatcher"
)

func TestOrMatcher(t *testing.T) {
	cases := []struct {
		pattern string
		input   string
		output  bool
	}{
		{
			pattern: "dv2ray.com",
			input:   "www.v2ray.com",
			output:  true,
		},
		{
			pattern: "dv2ray.com",
			input:   "v2ray.com",
			output:  true,
		},
		{
			pattern: "dv2ray.com",
			input:   "www.v3ray.com",
			output:  false,
		},
		{
			pattern: "dv2ray.com",
			input:   "2ray.com",
			output:  false,
		},
		{
			pattern: "dv2ray.com",
			input:   "xv2ray.com",
			output:  false,
		},
		{
			pattern: "fv2ray.com",
			input:   "v2ray.com",
			output:  true,
		},
		{
			pattern: "fv2ray.com",
			input:   "xv2ray.com",
			output:  false,
		},
		{
			pattern: "rv2ray.com",
			input:   "v2rayxcom",
			output:  true,
		},
		{
			pattern: "egeosite.dat:cn",
			input:   "www.baidu.com",
			output:  true,
		},
		{
			pattern: "egeosite.dat:cn",
			input:   "www.google.com",
			output:  false,
		},
		{
			pattern: "egeosite.dat:us",
			input:   "www.google.com",
			output:  true,
		},
	}
	external := map[string][]string{"geosite.dat:cn": []string{"dbaidu.com"}, "geosite.dat:us": []string{"dgoogle.com"}}
	for _, test := range cases {
		om := NewOrMatcher()
		err := om.ParsePattern(test.pattern, external)
		common.Must(err)
		if m := om.Match(test.input); m != test.output {
			t.Error("unexpected output: ", m, " for test case ", test)
		}
	}
}
