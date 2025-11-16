package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listOpsItemRelatedItemsCmd = &cobra.Command{
	Use:   "list-ops-item-related-items",
	Short: "Lists all related-item resources associated with a Systems Manager OpsCenter OpsItem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listOpsItemRelatedItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listOpsItemRelatedItemsCmd).Standalone()

		ssm_listOpsItemRelatedItemsCmd.Flags().String("filters", "", "One or more OpsItem filters.")
		ssm_listOpsItemRelatedItemsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listOpsItemRelatedItemsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_listOpsItemRelatedItemsCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem for which you want to list all related-item resources.")
	})
	ssmCmd.AddCommand(ssm_listOpsItemRelatedItemsCmd)
}
