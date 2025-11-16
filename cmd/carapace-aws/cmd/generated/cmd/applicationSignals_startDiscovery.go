package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_startDiscoveryCmd = &cobra.Command{
	Use:   "start-discovery",
	Short: "Enables this Amazon Web Services account to be able to use CloudWatch Application Signals by creating the *AWSServiceRoleForCloudWatchApplicationSignals* service-linked role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_startDiscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_startDiscoveryCmd).Standalone()

	})
	applicationSignalsCmd.AddCommand(applicationSignals_startDiscoveryCmd)
}
