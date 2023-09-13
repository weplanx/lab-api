package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/server/model"
	"testing"
)

func TestSetupImessage(t *testing.T) {
	ctx := context.TODO()
	err := model.SetupImessage(ctx, x.Db)
	assert.NoError(t, err)
}
