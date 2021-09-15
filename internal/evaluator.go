package internal

import (
	"amnesia-db/constants"
	"strings"
)

var GlobalKeyStore ConcurrentHashMap

func Eval(query string, pipe chan string) {
	queryExt := strings.Split(query, ":")
	var q = ParseFromString(queryExt[1])

	resp := ""
	if q.Op == GET {
		resp = string(GlobalKeyStore.Get(q.Key))
	} else if q.Op == SET {
		GlobalKeyStore.Insert(q.Key, q.Value, q.Options)
		resp = constants.OK
	} else {
		panic("Invalid Op Code Evaluated >> Panicking")
	}

	pipe <- queryExt[0] + ":" + resp
}
