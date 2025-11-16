package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getAssetTypeCmd = &cobra.Command{
	Use:   "get-asset-type",
	Short: "Gets an Amazon DataZone asset type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getAssetTypeCmd).Standalone()

	datazone_getAssetTypeCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the asset type exists.")
	datazone_getAssetTypeCmd.Flags().String("identifier", "", "The ID of the asset type.")
	datazone_getAssetTypeCmd.Flags().String("revision", "", "The revision of the asset type.")
	datazone_getAssetTypeCmd.MarkFlagRequired("domain-identifier")
	datazone_getAssetTypeCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getAssetTypeCmd)
}
