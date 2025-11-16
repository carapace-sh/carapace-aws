package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_listMonitorsCmd = &cobra.Command{
	Use:   "list-monitors",
	Short: "List all monitors in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_listMonitorsCmd).Standalone()

	networkflowmonitor_listMonitorsCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	networkflowmonitor_listMonitorsCmd.Flags().String("monitor-status", "", "The status of a monitor.")
	networkflowmonitor_listMonitorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_listMonitorsCmd)
}
