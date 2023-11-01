//go:generate weaver generate ./...

package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type app struct {
	weaver.Implements[weaver.Main]
}

func main() {}

func run(ctx context.Context, a *app) error {
	return nil
}
