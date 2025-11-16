package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listDeliverabilityTestReportsCmd = &cobra.Command{
	Use:   "list-deliverability-test-reports",
	Short: "Show a list of the predictive inbox placement tests that you've performed, regardless of their statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listDeliverabilityTestReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listDeliverabilityTestReportsCmd).Standalone()

		sesv2_listDeliverabilityTestReportsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListDeliverabilityTestReports` to indicate the position in the list of predictive inbox placement tests.")
		sesv2_listDeliverabilityTestReportsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListDeliverabilityTestReports`.")
	})
	sesv2Cmd.AddCommand(sesv2_listDeliverabilityTestReportsCmd)
}
