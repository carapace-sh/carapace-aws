package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_listMonitorsCmd = &cobra.Command{
	Use:   "list-monitors",
	Short: "Lists all of your monitors for Amazon CloudWatch Internet Monitor and their statuses, along with the Amazon Resource Name (ARN) and name of each monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_listMonitorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_listMonitorsCmd).Standalone()

		internetmonitor_listMonitorsCmd.Flags().Bool("include-linked-accounts", false, "A boolean option that you can set to `TRUE` to include monitors for linked accounts in a list of monitors, when you've set up cross-account sharing in Amazon CloudWatch Internet Monitor.")
		internetmonitor_listMonitorsCmd.Flags().String("max-results", "", "The number of monitor objects that you want to return with this call.")
		internetmonitor_listMonitorsCmd.Flags().String("monitor-status", "", "The status of a monitor.")
		internetmonitor_listMonitorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		internetmonitor_listMonitorsCmd.Flags().Bool("no-include-linked-accounts", false, "A boolean option that you can set to `TRUE` to include monitors for linked accounts in a list of monitors, when you've set up cross-account sharing in Amazon CloudWatch Internet Monitor.")
		internetmonitor_listMonitorsCmd.Flag("no-include-linked-accounts").Hidden = true
	})
	internetmonitorCmd.AddCommand(internetmonitor_listMonitorsCmd)
}
