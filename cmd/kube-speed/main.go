package main

import (
	"log"
	"github.com/docopt/docopt-go"
	"github.com/allen13/kube-speed/pkg/server"
	"github.com/allen13/kube-speed/pkg/job"
)

const version = "kube-speed 0.1.0"
const usage = `
Usage:
	kube-speed (job|server) [--completion-url=<url>] [--kube-speed-image=<image>]
	kube-speed job <request-id>
	kube-speed --help
	kube-speed --version

Options:
	--help                       Show this screen.
	--version                    Show version.
	--completion-url=<url>	     Where to send the job completion request to [default: http://kube-speed:1595/request].
	--kube-speed-image=<image>   kube-speed image for the job to use [default: allen13/kube-speed:latest].
`

func main() {
	// Parse args
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatalln(err)
	}
	completionURL := args["--completion-url"].(string)
	kubeSpeedImage :=  args["--kube-speed-image"].(string)

	if args["server"].(bool) {
		server.Start(completionURL, kubeSpeedImage)
	}

	if args["job"].(bool) {
		requestID := args["<request-id>"].(string)
		job.Run(completionURL, requestID)
	}
}