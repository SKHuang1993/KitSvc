package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "service"
	app.Version = "1.0"
	app.Usage = "starts the service daemon."
	app.Action = func(c *cli.Context) {
		server(c)
	}
	app.Flags = []cli.Flag{

		// Common flags.
		cli.StringFlag{
			EnvVar: "KITSVC_NAME",
			Name:   "name",
			Usage:  "the name of the service, exposed for service discovery.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_URL",
			Name:   "url",
			Usage:  "the url of the service.",
			Value:  "http://127.0.0.1:8080",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_ADDR",
			Name:   "addr",
			Usage:  "the address of the service (with the port).",
			Value:  "127.0.0.1:8080",
		},
		cli.IntFlag{
			EnvVar: "KITSVC_PORT",
			Name:   "port",
			Usage:  "the port of the service.",
			Value:  8080,
		},
		cli.StringFlag{
			EnvVar: "KITSVC_USAGE",
			Name:   "usage",
			Usage:  "the usage of the service, exposed for service discovery.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_JWT_SECRET",
			Name:   "jwt-secret",
			Usage:  "the secert used to encode the json web token.",
		},
		//cli.StringFlag{
		//	EnvVar: "KITSVC_VERSION",
		//	Name:   "version",
		//	Usage:  "the version of the service.",
		//},

		// Database flags.
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_NAME",
			Name:   "database-name",
			Usage:  "the name of the database.",
			Value:  "service",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_HOST",
			Name:   "database-host",
			Usage:  "the host of the database (with the port).",
			Value:  "127.0.0.1:3306",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_USER",
			Name:   "database-user",
			Usage:  "the user of the database.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_PASSWORD",
			Name:   "database-password",
			Usage:  "the password of the database.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_CHARSET",
			Name:   "database-charset",
			Usage:  "the charset of the database.",
			Value:  "utf8",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_DATABASE_LOC",
			Name:   "database-loc",
			Usage:  "the timezone of the database.",
			Value:  "Local",
		},
		cli.BoolFlag{
			EnvVar: "KITSVC_DATABASE_PARSE_TIME",
			Name:   "database-parse_time",
			Usage:  "parse the time.",
		},

		// Event store flags.
		cli.StringFlag{
			EnvVar: "KITSVC_ES_SERVER_URL",
			Name:   "es-url",
			Usage:  "the url of the event store server.",
			Value:  "http://127.0.0.1:2113",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_ES_USERNAME",
			Name:   "es-username",
			Usage:  "the username of the event store.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_ES_PASSWORD",
			Name:   "es-password",
			Usage:  "the password of the event store.",
		},

		// Prometheus flags.
		cli.StringFlag{
			EnvVar: "KITSVC_PROMETHEUS_NAMESPACE",
			Name:   "prometheus-namespace",
			Usage:  "the prometheus namespace.",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_PROMETHEUS_SUBSYSTEM",
			Name:   "prometheus-subsystem",
			Usage:  "the subsystem of the promethues.",
		},

		// Consul flags.
		cli.StringFlag{
			EnvVar: "KITSVC_CONSUL_CHECK_INTERVAL",
			Name:   "consul-check_interval",
			Usage:  "the interval of consul health check.",
			Value:  "10s",
		},
		cli.StringFlag{
			EnvVar: "KITSVC_CONSUL_CHECK_TIMEOUT",
			Name:   "consul-check_timeout",
			Usage:  "the timeout of consul health check.",
			Value:  "1s",
		},
		cli.StringSliceFlag{
			EnvVar: "KITSVC_CONSUL_TAGS",
			Name:   "consul-tags",
			Usage:  "the service tags for consul.",
			Value: &cli.StringSlice{
				"micro",
			},
		},
	}

	app.Run(os.Args)
}
