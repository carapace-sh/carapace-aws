package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetModelInterfaceRelationshipCmd = &cobra.Command{
	Use:   "describe-asset-model-interface-relationship",
	Short: "Retrieves information about an interface relationship between an asset model and an interface asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetModelInterfaceRelationshipCmd).Standalone()

	iotsitewise_describeAssetModelInterfaceRelationshipCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
	iotsitewise_describeAssetModelInterfaceRelationshipCmd.Flags().String("interface-asset-model-id", "", "The ID of the interface asset model.")
	iotsitewise_describeAssetModelInterfaceRelationshipCmd.MarkFlagRequired("asset-model-id")
	iotsitewise_describeAssetModelInterfaceRelationshipCmd.MarkFlagRequired("interface-asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetModelInterfaceRelationshipCmd)
}
