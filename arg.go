package go_args

import (
	"errors"
	"flag"
	"fmt"
)

func GetMapArgs(setError bool, keys ...string) (map[string]string, error) {

	var flags = map[string]*string{}

	for _, key := range keys {
		var value string
		flag.StringVar(&value, key, "", "")
		flags[key] = &value
	}

	flag.Parse()

	var resultFlags = map[string]string{}

	for k, v := range flags {

		if setError && *v == "" {
			return nil, errors.New(fmt.Sprintf("not found value for key: %s", k))
		}

		resultFlags[k] = *v
	}

	return resultFlags, nil
}
