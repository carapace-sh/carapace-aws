package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchAssociateProjectAssetsCmd = &cobra.Command{
	Use:   "batch-associate-project-assets",
	Short: "Associates a group (batch) of assets with an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchAssociateProjectAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_batchAssociateProjectAssetsCmd).Standalone()

		iotsitewise_batchAssociateProjectAssetsCmd.Flags().String("asset-ids", "", "The IDs of the assets to be associated to the project.")
		iotsitewise_batchAssociateProjectAssetsCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_batchAssociateProjectAssetsCmd.Flags().String("project-id", "", "The ID of the project to which to associate the assets.")
		iotsitewise_batchAssociateProjectAssetsCmd.MarkFlagRequired("asset-ids")
		iotsitewise_batchAssociateProjectAssetsCmd.MarkFlagRequired("project-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_batchAssociateProjectAssetsCmd)
}
