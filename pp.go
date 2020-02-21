package pp

import (
	"encoding/json"
	"fmt"
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

func NewJsonFormatter(opts... jsonFormatterOption) *JsonFormatter {
        return JsonFormatter{
                Prefix: jsonFormatterPrefixDefault,
                Indent: jsonFormatterIndentDefault,
        }
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
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "pp.Printjson: %s\n", err.Error())
	}
}
