package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listResourceComplianceSummariesCmd = &cobra.Command{
	Use:   "list-resource-compliance-summaries",
	Short: "Returns a resource-level summary count.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listResourceComplianceSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listResourceComplianceSummariesCmd).Standalone()

		ssm_listResourceComplianceSummariesCmd.Flags().String("filters", "", "One or more filters.")
		ssm_listResourceComplianceSummariesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listResourceComplianceSummariesCmd.Flags().String("next-token", "", "A token to start the list.")
	})
	ssmCmd.AddCommand(ssm_listResourceComplianceSummariesCmd)
}
