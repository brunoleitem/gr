package main

import (
	"github.com/brunoleitem/gr/cmd"
	"github.com/brunoleitem/gr/data"
)

func main() {
	data.OpenDb()
	cmd.Execute()
}
