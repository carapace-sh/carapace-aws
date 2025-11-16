package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteAssetCmd = &cobra.Command{
	Use:   "delete-asset",
	Short: "Deletes an asset in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteAssetCmd).Standalone()

	datazone_deleteAssetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the asset is deleted.")
	datazone_deleteAssetCmd.Flags().String("identifier", "", "The identifier of the asset that is deleted.")
	datazone_deleteAssetCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteAssetCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteAssetCmd)
}
