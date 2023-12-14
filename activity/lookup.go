package activity

import (
	"github.com/advanced-go/core/runtime"
)

var (
	overrideLookup func(string) string
)

func setOverrideLookup(t any) {
	if t == nil {
		overrideLookup = nil
		return
	}
	overrideLookup = runtime.OverrideLookup(t)
}

func lookup(key string) (string, bool) {
	if len(key) == 0 {
		return "", false
	}
	if overrideLookup != nil {
		val := overrideLookup(key)
		if len(val) > 0 {
			return val, true
		}
	}
	return "", false
}
