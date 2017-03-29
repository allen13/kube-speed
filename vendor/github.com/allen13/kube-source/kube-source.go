package main

import (
	"log"
	"github.com/docopt/docopt-go"
	"github.com/allen13/kube-source/app"
)

const version = "kube-source 0.1.0"
const usage = `
Usage:
	kube-source
	kube-source --help
	kube-source --version

Options:
	--help                       Show this screen.
	--version                    Show version.
`

func main() {
	// Parse args
	_, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.RunServer()
	if err != nil {
		log.Fatalln(err)
	}
}