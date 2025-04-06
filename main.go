package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/minnowo/traggo-sync/logging"
	"github.com/minnowo/traggo-sync/sync"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var TIME_FORMATS = []string{
	"02 Jan 2006 15:04 MST",
	"02 Jan 2006 15:04",

	"02 Jan 2006 MST",
	"02 Jan 2006",

	"2006-01-02 15:04:05 MST",
	"2006-01-02 15:04:05",

	"2006-01-02 MST",
	"2006-01-02",
}


func CLIDeleteMain(ctx context.Context, c *cli.Command) error {
	return fmt.Errorf("Comming soon")
}

func CLIExamplesMain(ctx context.Context, c *cli.Command) error {

	fmt.Println()
	fmt.Println()
	fmt.Println("Note: `ts` is the name of the executable, not part of the command\n")

	fmt.Println("Sync all missing tags:")
	fmt.Println("  ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-tags")

	fmt.Println()

	fmt.Println("Sync all missing tags and timespans:")
	fmt.Println("  ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-tags --missing-timespans")

	fmt.Println()

	fmt.Println("Sync missing timespans which overlap with `2025-01-05 EST` or later:")
	fmt.Println("  ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-timespans --time-start \"2025-01-05 EST\"")

	fmt.Println()

	fmt.Println("Sync missing timespans which overlap with `01 jan 2025 EST` until `03 jan 2025 EST` (inclusive):")
	fmt.Print("  ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-timespans --time-start \"01 jan 2025 EST\"")
	fmt.Println(" --time-end \"03 jan 2025 EST\"")

	fmt.Println()

	fmt.Println("Sync to more than one servers:")
	fmt.Println("  ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 -t http://localhost:3032 -t http://localhost:3033 --missing-tags")

	fmt.Println()

	fmt.Println("Sync with different credentials:")
	fmt.Println("  ts sync -i http://serverA.com -u adminA -p adminA -t http://serverB.com -u adminB -p adminB")

	fmt.Println()
	fmt.Println()

	return nil
}

func CLITimeMain(ctx context.Context, c *cli.Command) error {

	fmt.Println("Time formats supported for time ranges")
	fmt.Println()

	now := time.Now()

	for _, format := range TIME_FORMATS {

		fmt.Printf("Format: %s\n", format)
		fmt.Printf("Examples:\n")
		fmt.Printf("  %s\n", now.Format(format))

		for i := 1; i <= 2; i++ {
			daysAgo := rand.Intn(1000)
			randomTime := now.AddDate(0, 0, -daysAgo)

			// so there is never confusion about something like 2025-01-01 being year-m-d or year-d-m
			for randomTime.Day() < 13 {
				randomTime = randomTime.Add(time.Hour * 24)
			}

			dateStr := randomTime.Format(format)
			fmt.Printf("  %s", dateStr)

			parsedTime, err := time.Parse(format, dateStr)
			if err != nil {
				return err
			} else {
				fmt.Println("   ->   ", parsedTime)
			}
		}
		fmt.Println()
	}

	return nil
}

func CLISyncMain(c context.Context, cli *cli.Command) error {

	targetStrs := cli.Value("target").([]string)
	userStrs := cli.Value("user").([]string)
	passStrs := cli.Value("pass").([]string)

	useSameLogin := true

	if len(targetStrs) < 1 {
		return fmt.Errorf("must provide at least one target")
	}

	if len(userStrs) < 1 || len(passStrs) < 1 {
		return fmt.Errorf("must have at least one username and password")
	}

	if len(userStrs) != len(passStrs) {
		return fmt.Errorf("must have the same number of usernames as passwords")
	}

	if len(userStrs) > 1 || len(passStrs) > 1 {
		useSameLogin = false
		if len(targetStrs)+1 != len(userStrs) {
			return fmt.Errorf("must have the same number of usernames, passwords and hosts")
		}
	}

	if useSameLogin {
		log.Info().Msg("only 1 username / password given, using same login for all servers")
	}

	var run sync.RunConfig

	run.From.Host = cli.Value("source").(string)
	run.From.User = userStrs[0]
	run.From.Pass = passStrs[0]

	run.DryRun = cli.Value("dry-run").(bool)
	run.Dashboards = cli.Value("dashboards").(bool)

	run.AllTags = cli.Value("all-tags").(bool)
	run.MissingTags = cli.Value("missing-tags").(bool)

	run.MissingTimespans = cli.Value("missing-timespans").(bool)
	run.ReplaceTimespans = cli.Value("replace-timespans").(bool)
	run.DeleteTimespans = false

	run.StartTime = cli.Value("time-start").(time.Time)
	run.EndTime = cli.Value("time-end").(time.Time)

	var j int
	if useSameLogin {
		j = 0
	} else {
		j = 1
	}

	for i := 0; i < len(targetStrs); i++ {

		run.To = append(run.To, sync.TraggoServer{
			Host: targetStrs[i],
			User: userStrs[j],
			Pass: passStrs[j],
		})
		if !useSameLogin {
			j++
		}
	}

	return sync.SyncFrom(run)
}

func main() {

	logging.InitFromEnv()

	cmd := &cli.Command{
        Name: "Traggo Sync",
        // Description: "A tool for syncing data between Traggo servers",
        Usage: "A tool for syncing data between Traggo servers",
		Commands: []*cli.Command{
			{
				Name:        "examples",
				Description: "Shows some examples of the sync command",
				Action:      CLIExamplesMain,
			},
			{
				Name:        "time",
				Description: "Shows the various time formats for the --time-start and --time-end flags with examples.",
				Action:      CLITimeMain,
			},
			{
				Name:        "sync",
				Description: "Sync from one Traggo server to another",
				Action:      CLISyncMain,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "source",
						Aliases:  []string{"s", "i"},
						Usage:    "The `HOST` of the source Traggo server. Will sync from this server to the target servers. Example: http://localhost:3030",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:     "target",
						Aliases:  []string{"t", "o"},
						Usage:    "The `HOST` of the target Traggo server. Will sync from the source server to this server. Example: http://localhost:3031",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:     "user",
						Aliases:  []string{"u"},
						Usage:    "The `USERNAME` to login as. The first provided username is used for the source server; Using only a single username will use the same user for all servers.",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:     "pass",
						Aliases:  []string{"p"},
						Usage:    "The `PASSWORD` to login with. Each username must have a password.",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "missing-tags",
						Usage:    "Sync tags from the source to the target that do not exist on the target.",
						Value:    false,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "all-tags",
						Usage:    "Sync tags from the source to the target, if the tag exists the color will be updated.",
						Value:    false,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "missing-timespans",
						Usage:    "Sync timespans from the source to the target that do not exist on the target.",
						Value:    false,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "replace-timespans",
						Usage:    "Sync timespans from the source to the target deleting timestamps on the target which are not on the source.",
						Value:    false,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "dry-run",
						Aliases:  []string{"n"},
						Usage:    "Simulate the run without making any changes to any Traggo server.",
						Value:    false,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "dashboards",
						Usage:    "Simulate the run without making any changes to any Traggo server.",
						Value:    false,
						Required: false,
					},
					&cli.TimestampFlag{
						Name:  "time-start",
						Usage: "Starting range for timespan sync. Sync all timespans which fall between (inclusive) this value and time-end",
						Value: sync.MIN_TIME,
						Config: cli.TimestampConfig{
							Layouts: TIME_FORMATS,
						},
						Required: false,
					},
					&cli.TimestampFlag{
						Name:     "time-end",
						Usage:    "Ending range for timespan sync. Sync all timespans which fall between (inclusive) time-start and this value.",
						Value:    sync.MAX_TIME,
						Required: false,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
