package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_updateAssetCmd = &cobra.Command{
	Use:   "update-asset",
	Short: "This operation updates an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_updateAssetCmd).Standalone()

	dataexchange_updateAssetCmd.Flags().String("asset-id", "", "The unique identifier for an asset.")
	dataexchange_updateAssetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
	dataexchange_updateAssetCmd.Flags().String("name", "", "The name of the asset.")
	dataexchange_updateAssetCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
	dataexchange_updateAssetCmd.MarkFlagRequired("asset-id")
	dataexchange_updateAssetCmd.MarkFlagRequired("data-set-id")
	dataexchange_updateAssetCmd.MarkFlagRequired("name")
	dataexchange_updateAssetCmd.MarkFlagRequired("revision-id")
	dataexchangeCmd.AddCommand(dataexchange_updateAssetCmd)
}
