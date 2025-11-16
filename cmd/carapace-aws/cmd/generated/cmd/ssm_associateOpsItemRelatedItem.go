package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_associateOpsItemRelatedItemCmd = &cobra.Command{
	Use:   "associate-ops-item-related-item",
	Short: "Associates a related item to a Systems Manager OpsCenter OpsItem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_associateOpsItemRelatedItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_associateOpsItemRelatedItemCmd).Standalone()

		ssm_associateOpsItemRelatedItemCmd.Flags().String("association-type", "", "The type of association that you want to create between an OpsItem and a resource.")
		ssm_associateOpsItemRelatedItemCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem to which you want to associate a resource as a related item.")
		ssm_associateOpsItemRelatedItemCmd.Flags().String("resource-type", "", "The type of resource that you want to associate with an OpsItem.")
		ssm_associateOpsItemRelatedItemCmd.Flags().String("resource-uri", "", "The Amazon Resource Name (ARN) of the Amazon Web Services resource that you want to associate with the OpsItem.")
		ssm_associateOpsItemRelatedItemCmd.MarkFlagRequired("association-type")
		ssm_associateOpsItemRelatedItemCmd.MarkFlagRequired("ops-item-id")
		ssm_associateOpsItemRelatedItemCmd.MarkFlagRequired("resource-type")
		ssm_associateOpsItemRelatedItemCmd.MarkFlagRequired("resource-uri")
	})
	ssmCmd.AddCommand(ssm_associateOpsItemRelatedItemCmd)
}
