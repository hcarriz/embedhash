package valid

import "embed"

//go:embed test.md uhoh/*
var _ embed.FS
