package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeLoggingOptionsCmd = &cobra.Command{
	Use:   "describe-logging-options",
	Short: "Retrieves the current IoT SiteWise logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeLoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeLoggingOptionsCmd).Standalone()

	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeLoggingOptionsCmd)
}
