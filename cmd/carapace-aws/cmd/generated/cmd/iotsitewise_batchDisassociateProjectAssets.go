package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchDisassociateProjectAssetsCmd = &cobra.Command{
	Use:   "batch-disassociate-project-assets",
	Short: "Disassociates a group (batch) of assets from an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchDisassociateProjectAssetsCmd).Standalone()

	iotsitewise_batchDisassociateProjectAssetsCmd.Flags().String("asset-ids", "", "The IDs of the assets to be disassociated from the project.")
	iotsitewise_batchDisassociateProjectAssetsCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_batchDisassociateProjectAssetsCmd.Flags().String("project-id", "", "The ID of the project from which to disassociate the assets.")
	iotsitewise_batchDisassociateProjectAssetsCmd.MarkFlagRequired("asset-ids")
	iotsitewise_batchDisassociateProjectAssetsCmd.MarkFlagRequired("project-id")
	iotsitewiseCmd.AddCommand(iotsitewise_batchDisassociateProjectAssetsCmd)
}
