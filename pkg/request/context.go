package request

import (
	"context"
	"time"

	"github.com/golangci/golangci-shared/pkg/logutil"
)

type Context struct {
	Ctx  context.Context
	Log  logutil.Log
	Lctx logutil.Context

	StartedAt time.Time
}