package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteAssetTypeCmd = &cobra.Command{
	Use:   "delete-asset-type",
	Short: "Deletes an asset type in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteAssetTypeCmd).Standalone()

	datazone_deleteAssetTypeCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the asset type is deleted.")
	datazone_deleteAssetTypeCmd.Flags().String("identifier", "", "The identifier of the asset type that is deleted.")
	datazone_deleteAssetTypeCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteAssetTypeCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteAssetTypeCmd)
}
