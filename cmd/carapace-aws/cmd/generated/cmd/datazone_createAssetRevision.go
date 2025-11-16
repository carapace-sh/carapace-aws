package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createAssetRevisionCmd = &cobra.Command{
	Use:   "create-asset-revision",
	Short: "Creates a revision of the asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createAssetRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createAssetRevisionCmd).Standalone()

		datazone_createAssetRevisionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createAssetRevisionCmd.Flags().String("description", "", "The revised description of the asset.")
		datazone_createAssetRevisionCmd.Flags().String("domain-identifier", "", "The unique identifier of the domain where the asset is being revised.")
		datazone_createAssetRevisionCmd.Flags().String("forms-input", "", "The metadata forms to be attached to the asset as part of asset revision.")
		datazone_createAssetRevisionCmd.Flags().String("glossary-terms", "", "The glossary terms to be attached to the asset as part of asset revision.")
		datazone_createAssetRevisionCmd.Flags().String("identifier", "", "The identifier of the asset.")
		datazone_createAssetRevisionCmd.Flags().String("name", "", "Te revised name of the asset.")
		datazone_createAssetRevisionCmd.Flags().String("prediction-configuration", "", "The configuration of the automatically generated business-friendly metadata for the asset.")
		datazone_createAssetRevisionCmd.Flags().String("type-revision", "", "The revision type of the asset.")
		datazone_createAssetRevisionCmd.MarkFlagRequired("domain-identifier")
		datazone_createAssetRevisionCmd.MarkFlagRequired("identifier")
		datazone_createAssetRevisionCmd.MarkFlagRequired("name")
	})
	datazoneCmd.AddCommand(datazone_createAssetRevisionCmd)
}
