package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_putLoggingOptionsCmd = &cobra.Command{
	Use:   "put-logging-options",
	Short: "Sets logging options for IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_putLoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_putLoggingOptionsCmd).Standalone()

		iotsitewise_putLoggingOptionsCmd.Flags().String("logging-options", "", "The logging options to set.")
		iotsitewise_putLoggingOptionsCmd.MarkFlagRequired("logging-options")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_putLoggingOptionsCmd)
}
