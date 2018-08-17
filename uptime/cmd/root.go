package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	uptime "github.com/toversus/go-unixlike-command/uptime/uptm"
)

var rootCmd = &cobra.Command{
	Use:   "uptime",
	Short: "uptime - Tell how long the system has been running.",
	Long: ` uptime gives a one line display of the following information.  The current time, how long the
system has been running, how many users are currently logged on, and the system load averages
for the past 1, 5, and 15 minutes.`,
	Run: func(cmd *cobra.Command, args []string) {
		uptm, err := uptime.New()
		if err != nil {
			fmt.Fprintf(os.Stderr, "uptime initialization error: %s", err)
			os.Exit(1)
		}

		if pretty {
			uptm.PrettyPrint()
			return
		}

		if since {
			uptm.SincePrint()
			return
		}

		uptm.Print()
	},
}

var (
	pretty, since bool
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "show uptime in pretty format")
	rootCmd.Flags().BoolVarP(&since, "since", "s", false, "system up since")
}
