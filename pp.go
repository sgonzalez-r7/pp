package pp

import (
	"encoding/json"
	"fmt"
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
