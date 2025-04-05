package main

import (
	"context"
	"fmt"
	"os"

	"github.com/minnowo/traggo-sync/logging"
	"github.com/minnowo/traggo-sync/sync"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func DoSync(c context.Context, run sync.RunConfig) error {

	if err := run.Validate(); err != nil {
		return err
	}

	return sync.SyncFrom(run)
}

func CLIMain(c context.Context, cli *cli.Command) error {

	targetStrs := cli.Value("target").([]string)
	userStrs := cli.Value("user").([]string)
	passStrs := cli.Value("pass").([]string)

	if len(targetStrs)+1 != len(userStrs) || len(userStrs) != len(passStrs) {
		return fmt.Errorf("must have the same number of usernames, passwords and hosts")
	}

	var run sync.RunConfig

	run.From.Host = cli.Value("source").(string)
	run.From.User = userStrs[0]
	run.From.Pass = passStrs[0]

	for i := 0; i < len(targetStrs); i++ {

		run.To = append(run.To, sync.TraggoServer{
			Host: targetStrs[i],
			User: userStrs[i+1],
			Pass: passStrs[i+1],
		})
	}

	return DoSync(c, run)
}

func main() {

	logging.InitFromEnv()

	cmd := &cli.Command{
		Action: CLIMain,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s", "i"},
				Usage:    "The `HOST` of the source Traggo server. Copy from this to the target servers.",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "target",
				Aliases:  []string{"t", "o"},
				Usage:    "The `HOST` of the target Traggo server. Copy from the source to the this server.",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "user",
				Aliases:  []string{"u"},
				Usage:    "The `USERNAME` ",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "pass",
				Aliases:  []string{"p"},
				Usage:    "The `PASSWORD` ",
				Required: true,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
