package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_disassociateAssetsCmd = &cobra.Command{
	Use:   "disassociate-assets",
	Short: "Disassociates a child asset from the given parent asset through a hierarchy defined in the parent asset's model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_disassociateAssetsCmd).Standalone()

	iotsitewise_disassociateAssetsCmd.Flags().String("asset-id", "", "The ID of the parent asset from which to disassociate the child asset.")
	iotsitewise_disassociateAssetsCmd.Flags().String("child-asset-id", "", "The ID of the child asset to disassociate.")
	iotsitewise_disassociateAssetsCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_disassociateAssetsCmd.Flags().String("hierarchy-id", "", "The ID of a hierarchy in the parent asset's model.")
	iotsitewise_disassociateAssetsCmd.MarkFlagRequired("asset-id")
	iotsitewise_disassociateAssetsCmd.MarkFlagRequired("child-asset-id")
	iotsitewise_disassociateAssetsCmd.MarkFlagRequired("hierarchy-id")
	iotsitewiseCmd.AddCommand(iotsitewise_disassociateAssetsCmd)
}
