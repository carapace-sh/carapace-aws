package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_listAppMonitorsCmd = &cobra.Command{
	Use:   "list-app-monitors",
	Short: "Returns a list of the Amazon CloudWatch RUM app monitors in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_listAppMonitorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_listAppMonitorsCmd).Standalone()

		rum_listAppMonitorsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		rum_listAppMonitorsCmd.Flags().String("next-token", "", "Use the token returned by the previous operation to request the next page of results.")
	})
	rumCmd.AddCommand(rum_listAppMonitorsCmd)
}
