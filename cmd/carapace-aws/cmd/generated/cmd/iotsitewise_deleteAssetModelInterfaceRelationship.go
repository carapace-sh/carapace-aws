package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteAssetModelInterfaceRelationshipCmd = &cobra.Command{
	Use:   "delete-asset-model-interface-relationship",
	Short: "Deletes an interface relationship between an asset model and an interface asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteAssetModelInterfaceRelationshipCmd).Standalone()

	iotsitewise_deleteAssetModelInterfaceRelationshipCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
	iotsitewise_deleteAssetModelInterfaceRelationshipCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_deleteAssetModelInterfaceRelationshipCmd.Flags().String("interface-asset-model-id", "", "The ID of the interface asset model.")
	iotsitewise_deleteAssetModelInterfaceRelationshipCmd.MarkFlagRequired("asset-model-id")
	iotsitewise_deleteAssetModelInterfaceRelationshipCmd.MarkFlagRequired("interface-asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_deleteAssetModelInterfaceRelationshipCmd)
}
