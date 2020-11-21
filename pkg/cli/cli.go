package cli

import (
	"errors"
	"flag"
	"os"
	"wifi-connection-analyzer/pkg/types"
)

func GetCLIArgs() (types.Options, error) {
	// check for length of args
	if len(os.Args) < 2 {
		return types.Options{}, errors.New("an output filepath argument is required")
	}

	// set flag arguments
	intervalSeconds := flag.Int("i", 30, "Interval at which the daemon will run (in seconds)")
	healthCheckEndpoint := flag.String("h", "http://clients3.google.com/generate_204", "The endpoint that will be called to check network status")

	// parse all flags
	flag.Parse()

	// parse the only Arg (non-flag argument)
	fileLocation := flag.Arg(0)

	var opts = types.Options{
		IntervalSeconds: *intervalSeconds,
		HealthCheckEndpoint: *healthCheckEndpoint,
		FileLocation: fileLocation,
	}

	return opts, nil
}
