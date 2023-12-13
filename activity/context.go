package activity

import "context"

type variantT struct{}

var (
	variantKey = variantT{}
)

// newVariantContext - creates a new Context with a variant
func newVariantContext(ctx context.Context, variant string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if len(variant) == 0 {
		return ctx
	}
	return context.WithValue(ctx, variantKey, variant)
}

// variantFromContext - return a variant from a Context
func variantFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	i := ctx.Value(variantKey)
	if i == nil {
		return "", false
	}
	if s, ok := i.(string); ok {
		return s, true
	}
	return "", false
}
