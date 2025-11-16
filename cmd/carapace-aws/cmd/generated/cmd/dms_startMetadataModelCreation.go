package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelCreationCmd = &cobra.Command{
	Use:   "start-metadata-model-creation",
	Short: "Creates source metadata model of the given type with the specified properties for schema conversion operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelCreationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startMetadataModelCreationCmd).Standalone()

		dms_startMetadataModelCreationCmd.Flags().String("metadata-model-name", "", "The name of the metadata model.")
		dms_startMetadataModelCreationCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_startMetadataModelCreationCmd.Flags().String("properties", "", "The properties of metadata model in JSON format.")
		dms_startMetadataModelCreationCmd.Flags().String("selection-rules", "", "The JSON string that specifies the location where the metadata model will be created.")
		dms_startMetadataModelCreationCmd.MarkFlagRequired("metadata-model-name")
		dms_startMetadataModelCreationCmd.MarkFlagRequired("migration-project-identifier")
		dms_startMetadataModelCreationCmd.MarkFlagRequired("properties")
		dms_startMetadataModelCreationCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_startMetadataModelCreationCmd)
}
