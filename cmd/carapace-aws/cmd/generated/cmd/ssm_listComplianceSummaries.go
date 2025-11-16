package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listComplianceSummariesCmd = &cobra.Command{
	Use:   "list-compliance-summaries",
	Short: "Returns a summary count of compliant and non-compliant resources for a compliance type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listComplianceSummariesCmd).Standalone()

	ssm_listComplianceSummariesCmd.Flags().String("filters", "", "One or more compliance or inventory filters.")
	ssm_listComplianceSummariesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_listComplianceSummariesCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmCmd.AddCommand(ssm_listComplianceSummariesCmd)
}
