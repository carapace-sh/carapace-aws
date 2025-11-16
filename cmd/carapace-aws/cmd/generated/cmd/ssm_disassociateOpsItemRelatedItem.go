package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_disassociateOpsItemRelatedItemCmd = &cobra.Command{
	Use:   "disassociate-ops-item-related-item",
	Short: "Deletes the association between an OpsItem and a related item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_disassociateOpsItemRelatedItemCmd).Standalone()

	ssm_disassociateOpsItemRelatedItemCmd.Flags().String("association-id", "", "The ID of the association for which you want to delete an association between the OpsItem and a related item.")
	ssm_disassociateOpsItemRelatedItemCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem for which you want to delete an association between the OpsItem and a related item.")
	ssm_disassociateOpsItemRelatedItemCmd.MarkFlagRequired("association-id")
	ssm_disassociateOpsItemRelatedItemCmd.MarkFlagRequired("ops-item-id")
	ssmCmd.AddCommand(ssm_disassociateOpsItemRelatedItemCmd)
}
