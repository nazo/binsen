package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("server", nil, func(c *do.Context) {
		// rebuilds and restarts when a watched file changes
		c.Start("main.go s", do.M{"$in": "./cmd"})
	}).Src("*.go", "**/*.go").
		Debounce(3000)
}

func main() {
	do.Godo(tasks)
}
