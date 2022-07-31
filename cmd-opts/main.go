package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"io/ioutil"
	"log"
)

type Context struct {
	Debug bool
}

type LsCmd struct {
	Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
}

func (l *LsCmd) Run(ctx *Context) error {
	fmt.Println("ls", l.Paths)
	for _, path := range l.Paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println("Name:", file.Name(), "\n\tIs Directory? ", file.IsDir())
		}
	}
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	// Rm RmCmd `cmd help:"Remove files."`
	Ls LsCmd `cmd help:"List paths."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
