package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_associateAssetsCmd = &cobra.Command{
	Use:   "associate-assets",
	Short: "Associates a child asset with the given parent asset through a hierarchy defined in the parent asset's model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_associateAssetsCmd).Standalone()

	iotsitewise_associateAssetsCmd.Flags().String("asset-id", "", "The ID of the parent asset.")
	iotsitewise_associateAssetsCmd.Flags().String("child-asset-id", "", "The ID of the child asset to be associated.")
	iotsitewise_associateAssetsCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_associateAssetsCmd.Flags().String("hierarchy-id", "", "The ID of a hierarchy in the parent asset's model.")
	iotsitewise_associateAssetsCmd.MarkFlagRequired("asset-id")
	iotsitewise_associateAssetsCmd.MarkFlagRequired("child-asset-id")
	iotsitewise_associateAssetsCmd.MarkFlagRequired("hierarchy-id")
	iotsitewiseCmd.AddCommand(iotsitewise_associateAssetsCmd)
}
