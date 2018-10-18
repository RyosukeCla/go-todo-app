package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("build", nil, func(c *do.Context) {
		c.Run("go build", do.M{"$in": "."})
	})

	p.Task("dev", nil, func(c *do.Context) {
		c.Run("go build", do.M{"$in": "."})
		c.Start("./go-todo-app")
	}).Src("*.go", "**/*.go").Debounce(1000)
}

func main() {
	do.Godo(tasks)
}
