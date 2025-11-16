package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getAssetCmd = &cobra.Command{
	Use:   "get-asset",
	Short: "This operation returns information about an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_getAssetCmd).Standalone()

		dataexchange_getAssetCmd.Flags().String("asset-id", "", "The unique identifier for an asset.")
		dataexchange_getAssetCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_getAssetCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_getAssetCmd.MarkFlagRequired("asset-id")
		dataexchange_getAssetCmd.MarkFlagRequired("data-set-id")
		dataexchange_getAssetCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_getAssetCmd)
}
