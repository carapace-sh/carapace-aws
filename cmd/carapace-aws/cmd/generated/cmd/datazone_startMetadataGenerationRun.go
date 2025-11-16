package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_startMetadataGenerationRunCmd = &cobra.Command{
	Use:   "start-metadata-generation-run",
	Short: "Starts the metadata generation run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_startMetadataGenerationRunCmd).Standalone()

	datazone_startMetadataGenerationRunCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	datazone_startMetadataGenerationRunCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where you want to start a metadata generation run.")
	datazone_startMetadataGenerationRunCmd.Flags().String("owning-project-identifier", "", "The ID of the project that owns the asset for which you want to start a metadata generation run.")
	datazone_startMetadataGenerationRunCmd.Flags().String("target", "", "The asset for which you want to start a metadata generation run.")
	datazone_startMetadataGenerationRunCmd.Flags().String("type", "", "The type of the metadata generation run.")
	datazone_startMetadataGenerationRunCmd.MarkFlagRequired("domain-identifier")
	datazone_startMetadataGenerationRunCmd.MarkFlagRequired("owning-project-identifier")
	datazone_startMetadataGenerationRunCmd.MarkFlagRequired("target")
	datazone_startMetadataGenerationRunCmd.MarkFlagRequired("type")
	datazoneCmd.AddCommand(datazone_startMetadataGenerationRunCmd)
}
