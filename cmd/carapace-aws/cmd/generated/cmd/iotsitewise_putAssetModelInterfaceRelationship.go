package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_putAssetModelInterfaceRelationshipCmd = &cobra.Command{
	Use:   "put-asset-model-interface-relationship",
	Short: "Creates or updates an interface relationship between an asset model and an interface asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_putAssetModelInterfaceRelationshipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_putAssetModelInterfaceRelationshipCmd).Standalone()

		iotsitewise_putAssetModelInterfaceRelationshipCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.Flags().String("interface-asset-model-id", "", "The ID of the interface asset model.")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.Flags().String("property-mapping-configuration", "", "The configuration for mapping properties from the interface asset model to the asset model where the interface is applied.")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.MarkFlagRequired("asset-model-id")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.MarkFlagRequired("interface-asset-model-id")
		iotsitewise_putAssetModelInterfaceRelationshipCmd.MarkFlagRequired("property-mapping-configuration")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_putAssetModelInterfaceRelationshipCmd)
}
