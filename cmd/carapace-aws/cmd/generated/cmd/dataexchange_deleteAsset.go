package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_deleteAssetCmd = &cobra.Command{
	Use:   "delete-asset",
	Short: "This operation deletes an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_deleteAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_deleteAssetCmd).Standalone()

		dataexchange_deleteAssetCmd.Flags().String("asset-id", "", "The unique identifier for an asset.")
		dataexchange_deleteAssetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_deleteAssetCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_deleteAssetCmd.MarkFlagRequired("asset-id")
		dataexchange_deleteAssetCmd.MarkFlagRequired("data-set-id")
		dataexchange_deleteAssetCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_deleteAssetCmd)
}
