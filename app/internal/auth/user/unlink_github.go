package user

import (
	"fmt"

	"github.com/golangci/golangci-api/app/internal/db"
	"github.com/golangci/golib/server/context"
	"github.com/golangci/golib/server/handlers/herrors"
)

func UnlinkGithub(ctx *context.C) error {
	ga, err := GetGithubAuth(ctx)
	if err != nil {
		return herrors.New(err, "can't get github auth")
	}

	if err = ga.Delete(db.Get(ctx).Unscoped()); err != nil {
		return fmt.Errorf("can't delete github auth: %s", err)
	}

	return nil
}