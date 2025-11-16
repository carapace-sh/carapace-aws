package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listComplianceItemsCmd = &cobra.Command{
	Use:   "list-compliance-items",
	Short: "For a specified resource ID, this API operation returns a list of compliance statuses for different resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listComplianceItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listComplianceItemsCmd).Standalone()

		ssm_listComplianceItemsCmd.Flags().String("filters", "", "One or more compliance filters.")
		ssm_listComplianceItemsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listComplianceItemsCmd.Flags().String("next-token", "", "A token to start the list.")
		ssm_listComplianceItemsCmd.Flags().String("resource-ids", "", "The ID for the resources from which to get compliance information.")
		ssm_listComplianceItemsCmd.Flags().String("resource-types", "", "The type of resource from which to get compliance information.")
	})
	ssmCmd.AddCommand(ssm_listComplianceItemsCmd)
}
