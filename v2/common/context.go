package common

import (
	"context"
	"time"
)

var NewDailContext = func() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
