package main

import (
	"log"

	"github.com/docopt/docopt-go"
	"github.com/jboursiquot/rfc"
)

func main() {
	usage := `rfc

	Usage:
		rfc update
		rfc save --path <path> --format <format>
		rfc serve --host <host> --port <port>
		rfc -h | --help
		rfc --version

	Options:
		--host <host> Server host.
		--port <port> Server port.
		--path <path> Output path.
		--format Output format (json, csv or bolt), default = json
		-h --help   Show this screen.
		--version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "RFC", false)
	log.Println(arguments)

	switch {
	case arguments["update"] == true:
		rfc.Update()
	case arguments["serve"] == true:
		// TODO serve
	case arguments["save"] == true:
		// TODO output
	}
}
