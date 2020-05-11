# pp

Convenience functions for pretty-printing in Go

# Basic Usage

`import "github.com/sgonzalez-r7/pp"`

### Use [SpewDump](https://github.com/davecgh/go-spew) to print `k-v` pairs to `STDERR`
`pp.SpewDump("k1", v1, "k2", v2)`

### Use [SpewDump](https://github.com/davecgh/go-spew) to write `k-v` pairs to a file
`
if out, err := os.Create("/path/to/file"); err == nil {
    pp.FspewDump("k1", v1, "k2", v2)
  } else {
    fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
  }
`

### print json to `STDERR`
`pp.Printjson(v)`

### write json to file
```
if out, err := os.Create("/path/to/file"); err == nil {
  pp.Fprintjson(out, templateData)
} else {
  fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
}

```
