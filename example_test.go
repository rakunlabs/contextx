package contextx_test

import (
	"context"
	"fmt"

	"github.com/rakunlabs/contextx"
)

func Example() {
	ctx := contextx.WithValue(context.Background(), "secret", "xxx")

	if v, ok := contextx.Value[string](ctx, "secret"); ok {
		fmt.Println("[" + v + "]")
	}

	ctx = contextx.WithValue(ctx, "another", "yyy", contextx.WithKey("customKey"))

	if v, ok := contextx.Value[string](ctx, "another"); ok {
		fmt.Println("[" + v + "]")
	} else {
		// expected not found
		fmt.Println("not found")
	}

	if v, ok := contextx.Value[string](ctx, "secret", contextx.WithKey("customKey")); ok {
		fmt.Println("[" + v + "]")
	} else {
		// expected not found
		fmt.Println("not found")
	}

	if v, ok := contextx.Value[string](ctx, "another", contextx.WithKey("customKey")); ok {
		// expected found
		fmt.Println("[" + v + "]")
	}

	// Output:
	// [xxx]
	// not found
	// not found
	// [yyy]
}
