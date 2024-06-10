// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://devniel93.io

package tui

import (
	whatChangedModel "github.com/devniel93/libopenapi/what-changed/model"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"testing"
)

// reported in https://github.com/devniel93/my-openapi-changes/issues/115
func TestBuildTreeModel_SliceOfChanges(t *testing.T) {

	// create a dummy slice of whatChangedModel slice
	changed := []*whatChangedModel.Change{
		{
			ChangeType: whatChangedModel.PropertyAdded,
		},
		{
			ChangeType: whatChangedModel.PropertyRemoved,
		},
	}

	p := &whatChangedModel.PropertyChanges{Changes: changed}
	f := tview.TreeNode{}
	v := buildTreeNode(&f, p)
	assert.NotNil(t, v)
	assert.Equal(t, 2, len(v.GetChildren()))
}
