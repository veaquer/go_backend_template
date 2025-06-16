package auth

import (
	"github.com/veaquer/go_backend_template/pkg/constants"
	"context"
	"errors"
)

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	id, ok := ctx.Value(constants.UserIdCtxKey).(uint)
	if !ok {
		return 0, errors.New("user id not found in context")
	}

	return id, nil
}
