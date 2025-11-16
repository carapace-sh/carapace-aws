package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteSharedTrustStoreAssociationCmd = &cobra.Command{
	Use:   "delete-shared-trust-store-association",
	Short: "Deletes a shared trust store association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteSharedTrustStoreAssociationCmd).Standalone()

	elbv2_deleteSharedTrustStoreAssociationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	elbv2_deleteSharedTrustStoreAssociationCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2_deleteSharedTrustStoreAssociationCmd.MarkFlagRequired("resource-arn")
	elbv2_deleteSharedTrustStoreAssociationCmd.MarkFlagRequired("trust-store-arn")
	elbv2Cmd.AddCommand(elbv2_deleteSharedTrustStoreAssociationCmd)
}
