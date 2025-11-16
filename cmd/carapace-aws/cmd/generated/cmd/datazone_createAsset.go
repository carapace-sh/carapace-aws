package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createAssetCmd = &cobra.Command{
	Use:   "create-asset",
	Short: "Creates an asset in Amazon DataZone catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createAssetCmd).Standalone()

	datazone_createAssetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_createAssetCmd.Flags().String("description", "", "Asset description.")
	datazone_createAssetCmd.Flags().String("domain-identifier", "", "Amazon DataZone domain where the asset is created.")
	datazone_createAssetCmd.Flags().String("external-identifier", "", "The external identifier of the asset.")
	datazone_createAssetCmd.Flags().String("forms-input", "", "Metadata forms attached to the asset.")
	datazone_createAssetCmd.Flags().String("glossary-terms", "", "Glossary terms attached to the asset.")
	datazone_createAssetCmd.Flags().String("name", "", "Asset name.")
	datazone_createAssetCmd.Flags().String("owning-project-identifier", "", "The unique identifier of the project that owns this asset.")
	datazone_createAssetCmd.Flags().String("prediction-configuration", "", "The configuration of the automatically generated business-friendly metadata for the asset.")
	datazone_createAssetCmd.Flags().String("type-identifier", "", "The unique identifier of this asset's type.")
	datazone_createAssetCmd.Flags().String("type-revision", "", "The revision of this asset's type.")
	datazone_createAssetCmd.MarkFlagRequired("domain-identifier")
	datazone_createAssetCmd.MarkFlagRequired("name")
	datazone_createAssetCmd.MarkFlagRequired("owning-project-identifier")
	datazone_createAssetCmd.MarkFlagRequired("type-identifier")
	datazoneCmd.AddCommand(datazone_createAssetCmd)
}
