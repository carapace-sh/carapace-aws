package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_cancelMetadataGenerationRunCmd = &cobra.Command{
	Use:   "cancel-metadata-generation-run",
	Short: "Cancels the metadata generation run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_cancelMetadataGenerationRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_cancelMetadataGenerationRunCmd).Standalone()

		datazone_cancelMetadataGenerationRunCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the metadata generation run is to be cancelled.")
		datazone_cancelMetadataGenerationRunCmd.Flags().String("identifier", "", "The ID of the metadata generation run.")
		datazone_cancelMetadataGenerationRunCmd.MarkFlagRequired("domain-identifier")
		datazone_cancelMetadataGenerationRunCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_cancelMetadataGenerationRunCmd)
}
