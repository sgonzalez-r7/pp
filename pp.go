package pp

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"os"
)

type JsonFormatter struct {
        Prefix string
        Indent string
}

const (
        jsonFormatterPrefixDefault = ""     // empty string
        jsonFormatterIndentDefault = "    " // 4 spaces
)

type jsonFormatterOption func(*JsonFormatter)

func Indent(s string) jsonFormatterOption {
	return func(f *JsonFormatter) {
		f.Indent = s
	}
}

func Prefix(s string) jsonFormatterOption {
	return func(f *JsonFormatter) {
		f.Prefix = s
	}
}

func NewJsonFormatter(opts... jsonFormatterOption) *JsonFormatter {
	f := &JsonFormatter{
		Prefix: jsonFormatterPrefixDefault,
		Indent: jsonFormatterIndentDefault,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *JsonFormatter) Fprintjson(w io.Writer, v interface{}) {
        enc := json.NewEncoder(w)
        enc.SetIndent(f.Prefix, f.Indent)
        if err := enc.Encode(v); err != nil {
                fmt.Fprintf(os.Stderr, "pp: %s\n", err.Error())
        }
}

func (f *JsonFormatter) Printjson(v interface{}) {
        f.Fprintjson(os.Stderr, v)
}

func Fprintjson(w io.Writer, v interface{}) {
        f := NewJsonFormatter()
        f.Fprintjson(w, v)
}

func Printjson(v interface{}) {
	Fprintjson(os.Stderr, v)
}

//
// Spew
//

var (
	scs spew.ConfigState = spew.ConfigState{
		Indent: "    ",
		DisableMethods: true,
		DisablePointerAddresses: true,
	}
)

type KVPair struct {
	K string
	V interface{}
}

func FspewDump(w io.Writer, args ...interface{}) {
	if len(args)%2 != 0 {
		msg := fmt.Sprintf("pp.FspewDump: kvpairs must be even. got %d", len(args))
		scs.Fdump(w, msg)
		return
	}

	kvps := make([]KVPair, len(args)/2)
	for i := 0; i < len(kvps); i++ {
		kvps[i] = KVPair{K: args[2*i].(string), V: args[2*i+1]}
	}

	scs.Fdump(w, kvps)
}

func SpewDump(args ...interface{}) {
	if len(args)%2 != 0 {
		msg := fmt.Sprintf("pp.FspewDump: kvpairs must be even. got %d", len(args))
		scs.Dump(msg)
		return
	}

	kvps := make([]KVPair, len(args)/2)
	for i := 0; i < len(kvps); i++ {
		kvps[i] = KVPair{K: args[2*i].(string), V: args[2*i+1]}
	}

	scs.Dump(kvps)
}
