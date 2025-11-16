package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_listMonitorsCmd = &cobra.Command{
	Use:   "list-monitors",
	Short: "Returns a list of all of your monitors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_listMonitorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_listMonitorsCmd).Standalone()

		networkmonitor_listMonitorsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		networkmonitor_listMonitorsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmonitor_listMonitorsCmd.Flags().String("state", "", "The list of all monitors and their states.")
	})
	networkmonitorCmd.AddCommand(networkmonitor_listMonitorsCmd)
}
