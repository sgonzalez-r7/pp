package pp

import (
	"encoding/json"
	"fmt"
	"os"
)

func Printjson(v interface{}) {
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "pp.Printjson: %s\n", err.Error())
	}
}
