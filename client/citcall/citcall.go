package citcall

import (
	"context"
)

type Client interface {
	GetCountries(ctx context.Context) (Countries, error)
}
