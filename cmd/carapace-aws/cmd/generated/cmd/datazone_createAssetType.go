package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createAssetTypeCmd = &cobra.Command{
	Use:   "create-asset-type",
	Short: "Creates a custom asset type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createAssetTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createAssetTypeCmd).Standalone()

		datazone_createAssetTypeCmd.Flags().String("description", "", "The descripton of the custom asset type.")
		datazone_createAssetTypeCmd.Flags().String("domain-identifier", "", "The unique identifier of the Amazon DataZone domain where the custom asset type is being created.")
		datazone_createAssetTypeCmd.Flags().String("forms-input", "", "The metadata forms that are to be attached to the custom asset type.")
		datazone_createAssetTypeCmd.Flags().String("name", "", "The name of the custom asset type.")
		datazone_createAssetTypeCmd.Flags().String("owning-project-identifier", "", "The identifier of the Amazon DataZone project that is to own the custom asset type.")
		datazone_createAssetTypeCmd.MarkFlagRequired("domain-identifier")
		datazone_createAssetTypeCmd.MarkFlagRequired("forms-input")
		datazone_createAssetTypeCmd.MarkFlagRequired("name")
		datazone_createAssetTypeCmd.MarkFlagRequired("owning-project-identifier")
	})
	datazoneCmd.AddCommand(datazone_createAssetTypeCmd)
}
