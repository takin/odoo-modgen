package main

import (
	"flag"
	"os"

	"github.com/takin/odoo-modgen/cmd"
)

func main() {
	curdir, _ := os.Getwd()
	workingDir := flag.String("path", curdir, "specify the working directory")
	odooVersion := flag.Int("odoo", 12, "Specify Odoo version")
	flag.Parse()

	o := &cmd.OdooModule{
		Name:        flag.Arg(0),
		Path:        *workingDir,
		OdooVersion: *odooVersion,
	}

	o.Generate()
}
