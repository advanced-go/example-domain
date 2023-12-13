package activity

import (
	"context"
)

type lookupT struct{}

var (
	lookupKey = lookupT{}
)

// NewLookupContext - creates a new Context with a lookup
func NewLookupContext(ctx context.Context, lookup any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if lookup == nil {
		return ctx
	}
	return context.WithValue(ctx, lookupKey, funcFromType(lookup))
}

// LookupFromContext - return a lookup from a Context
func LookupFromContext(ctx context.Context, value string) string {
	if ctx == nil {
		return value
	}
	i := ctx.Value(lookupKey)
	if i == nil {
		return value
	}
	if fn, ok := i.(func(string) string); ok {
		return fn(value)
	}
	return value
}

func funcFromType(t any) func(string) string {
	switch ptr := t.(type) {
	case string:
		return func(k string) string { return ptr }
	case map[string]string:
		return func(k string) string {
			v := ptr[k]
			if len(v) > 0 {
				return v
			}
			return k
		}
	case func(string) string:
		return ptr
	}
	return func(k string) string { return k }
}
