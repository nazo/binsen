// +build tools

package main

import (
	_ "github.com/volatiletech/sqlboiler"
	_ "gopkg.in/godo.v2/cmd/godo"
	_ "github.com/pressly/goose/cmd/goose"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"
)

