package xgb

import (
	_ "embed"
)

// This package only contains the model bytes
// This makes it possible to embed the model in the binary
// and allows for easier portable builds / testing

//go:embed xgb.model
var ModelBytes []byte
