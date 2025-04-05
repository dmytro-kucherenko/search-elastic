package common

import (
	"context"
)

type Versions map[string]uint

type Migration interface {
	Name() string
	Exists() bool
	Make(context.Context) error
}

type NewMigration func(*Tool) Migration
