//go:generate weaver generate ./...

package main

import (
	"context"
)

type app struct{}

func main() {}

func run(ctx context.Context, a *app) error {
	return nil
}
