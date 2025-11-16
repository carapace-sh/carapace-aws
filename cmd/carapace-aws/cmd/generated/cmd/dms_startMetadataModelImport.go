package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelImportCmd = &cobra.Command{
	Use:   "start-metadata-model-import",
	Short: "Loads the metadata for all the dependent database objects of the parent object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startMetadataModelImportCmd).Standalone()

		dms_startMetadataModelImportCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_startMetadataModelImportCmd.Flags().Bool("no-refresh", false, "If `true`, DMS loads metadata for the specified objects from the source database.")
		dms_startMetadataModelImportCmd.Flags().String("origin", "", "Whether to load metadata to the source or target database.")
		dms_startMetadataModelImportCmd.Flags().Bool("refresh", false, "If `true`, DMS loads metadata for the specified objects from the source database.")
		dms_startMetadataModelImportCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to import.")
		dms_startMetadataModelImportCmd.MarkFlagRequired("migration-project-identifier")
		dms_startMetadataModelImportCmd.Flag("no-refresh").Hidden = true
		dms_startMetadataModelImportCmd.MarkFlagRequired("origin")
		dms_startMetadataModelImportCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_startMetadataModelImportCmd)
}
