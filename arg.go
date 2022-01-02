package go_args

import "flag"

func GetMapArgs(keys ...string) map[string]string {

	var flags = map[string]*string{}

	for _, key := range keys {
		var value *string
		flag.StringVar(value, key, "", "")
		flags[key] = value
	}

	flag.Parse()

	var resultFlags = map[string]string{}

	for k, v := range flags {
		resultFlags[k] = *v
	}

	return resultFlags
}
