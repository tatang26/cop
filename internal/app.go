package internal

import (
	appAssets "cop/internal/assets"
	"cop/internal/templates"

	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/render"
)

var (
	renderMW = render.InCtx(templates.FS,
		render.WithDefaultLayout("layout.html"),
	)

	assetsManager = assets.NewManager(appAssets.FS)
)
