package sophos_test

import (
	"testing"

	"github.com/esurdam/go-sophos"
)

func TestErrors_Error(t *testing.T) {
	tests := []struct {
		name string
		ee   sophos.Errors
		want string
	}{
		{"testNoFatalErrors", sophos.Errors{
			{
				Oattrs: []string{
					"class",
					"type",
				},
				Class:     "packetfilter",
				Fatal:     0,
				Format:    "The %_O object requires %_d for the %_A attribute.",
				Msgtype:   "DATATYPE_OBJECT_ATTRIBUTE",
				Name:      "The group object requires a Perl array for the members list attribute.",
				NeverHide: 0,
				Type:      "group",
			}, {
				Oattrs: []string{
					"class",
					"type",
				},
				Class:     "packetfilter",
				Fatal:     0,
				Format:    "The %_O object requires %_d for the %_A attribute.",
				Msgtype:   "DATATYPE_OBJECT_ATTRIBUTE",
				Name:      "The group object requires a Perl array for the members list attribute.",
				NeverHide: 0,
				Type:      "group",
			},
		}, "client do: error from UTM server: The group object requires a Perl array for the members list attribute."},
		{"testFatalErrors", sophos.Errors{
			{
				Oattrs: []string{
					"class",
					"type",
				},
				Class:     "packetfilter",
				Fatal:     1,
				Format:    "The %_O object requires %_d for the %_A attribute.",
				Msgtype:   "DATATYPE_OBJECT_ATTRIBUTE",
				Name:      "The TEST object requires a TEST array for the members list attribute.",
				NeverHide: 0,
				Type:      "group",
			}}, "client do: FATAL error from UTM server: The TEST object requires a TEST array for the members list attribute."},
		{"testNoErrors", sophos.Errors{}, "error accessing UTM interface: check status code. No Errors were retuned in response body."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ee.Error(); got != tt.want {
				t.Errorf("Errors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFatalErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"testIsFatalErrorsT", args{sophos.Errors{{Fatal: 1}}}, true},
		{"testIsFatal*ErrorsT", args{&sophos.Errors{{Fatal: 1}}}, true},
		{"testIsFatalErrorT", args{sophos.Error{Fatal: 1}}, true},
		{"testIsFatal*ErrorT", args{&sophos.Error{Fatal: 1}}, true},
		{"testIsFatalErrorsT", args{sophos.Errors{{Fatal: 0}}}, false},
		{"testIsFatal*ErrorsT", args{&sophos.Errors{{Fatal: 0}}}, false},
		{"testIsFatalErrorT", args{sophos.Error{Fatal: 0}}, false},
		{"testIsFatal*ErrorT", args{&sophos.Error{Fatal: 0}}, false},
		{"testIsFatalnil", args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sophos.IsFatalErr(tt.args.err); got != tt.want {
				t.Errorf("IsFatalErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
