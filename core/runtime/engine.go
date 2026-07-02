package runtime

import "context"

type Engine interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}