package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getMetadataGenerationRunCmd = &cobra.Command{
	Use:   "get-metadata-generation-run",
	Short: "Gets a metadata generation run in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getMetadataGenerationRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getMetadataGenerationRunCmd).Standalone()

		datazone_getMetadataGenerationRunCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain the metadata generation run of which you want to get.")
		datazone_getMetadataGenerationRunCmd.Flags().String("identifier", "", "The identifier of the metadata generation run.")
		datazone_getMetadataGenerationRunCmd.MarkFlagRequired("domain-identifier")
		datazone_getMetadataGenerationRunCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getMetadataGenerationRunCmd)
}
