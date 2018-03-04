package repos

import (
	"fmt"
	"os"

	"github.com/golangci/golangci-api/app/internal/auth/user"
	"github.com/golangci/golangci-api/app/internal/db"
	"github.com/golangci/golangci-api/app/internal/github"
	"github.com/golangci/golangci-api/app/models"
	"github.com/golangci/golangci-api/app/utils"
	"github.com/golangci/golib/server/context"
	"github.com/golangci/golib/server/handlers/herrors"
	gh "github.com/google/go-github/github"
)

func DeactivateRepo(ctx *context.C, owner, repo string) (*models.GithubRepo, error) {
	gc, err := github.GetClient(ctx)
	if err != nil {
		return nil, herrors.New(err, "can't get github client")
	}

	var gr models.GithubRepo
	err = models.NewGithubRepoQuerySet(db.Get(ctx)).
		NameEq(fmt.Sprintf("%s/%s", owner, repo)).
		One(&gr)
	if err != nil {
		return nil, fmt.Errorf("can't get repo %s/%s: %s", owner, repo, err)
	}

	_, err = gc.Repositories.DeleteHook(ctx.Ctx, owner, repo, gr.GithubHookID)
	if err != nil {
		return nil, fmt.Errorf("can't delete hook %d from github repo %s/%s: %s",
			gr.GithubHookID, owner, repo, err)
	}

	if err = gr.Delete(db.Get(ctx)); err != nil {
		return nil, fmt.Errorf("can't delete github repo: %s", err)
	}

	return &gr, nil
}

func GetWebhookURLPathForRepo(name, hookID string) string {
	return fmt.Sprintf("/v1/repos/%s/hooks/%s", name, hookID)
}

func ActivateRepo(ctx *context.C, ga *models.GithubAuth, owner, repo string) (*models.GithubRepo, error) {
	gc, err := github.GetClient(ctx)
	if err != nil {
		return nil, herrors.New(err, "can't get github client")
	}

	hookID, err := utils.GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("can't generate hook id: %s", err)
	}

	gr := models.GithubRepo{
		UserID: ga.UserID,
		Name:   fmt.Sprintf("%s/%s", owner, repo),
		HookID: hookID,
	}

	web := "web"
	hookURL := os.Getenv("GITHUB_CALLBACK_HOST") + GetWebhookURLPathForRepo(gr.Name, gr.HookID)
	hook := gh.Hook{
		Name:   &web,
		Events: []string{"pull_request"},
		Config: map[string]interface{}{
			"url":          hookURL,
			"content_type": "json",
		},
	}
	rh, _, err := gc.Repositories.CreateHook(ctx.Ctx, owner, repo, &hook)
	if err != nil {
		return nil, fmt.Errorf("can't post hook %v to github: %s", hook, err)
	}

	gr.GithubHookID = rh.GetID()
	if err := gr.Create(db.Get(ctx)); err != nil {
		return nil, herrors.New(err, "can't create github repo")
	}

	return &gr, nil
}

func DeactivateAll(ctx *context.C) error {
	userID, err := user.GetCurrentID(ctx)
	if err != nil {
		return fmt.Errorf("can't get current user id: %s", err)
	}

	// TODO: remove all hooks
	err = models.NewGithubRepoQuerySet(db.Get(ctx)).
		UserIDEq(userID).Delete()
	if err != nil {
		return fmt.Errorf("can't delete all repos of user %d: %s", userID, err)
	}

	return nil
}

func ArePrivateReposEnabledForUser(_ *context.C) bool {
	return false
}