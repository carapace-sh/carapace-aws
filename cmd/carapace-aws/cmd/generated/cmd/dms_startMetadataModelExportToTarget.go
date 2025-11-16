package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelExportToTargetCmd = &cobra.Command{
	Use:   "start-metadata-model-export-to-target",
	Short: "Applies converted database objects to your target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelExportToTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startMetadataModelExportToTargetCmd).Standalone()

		dms_startMetadataModelExportToTargetCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_startMetadataModelExportToTargetCmd.Flags().String("overwrite-extension-pack", "", "Whether to overwrite the migration project extension pack.")
		dms_startMetadataModelExportToTargetCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to export.")
		dms_startMetadataModelExportToTargetCmd.MarkFlagRequired("migration-project-identifier")
		dms_startMetadataModelExportToTargetCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_startMetadataModelExportToTargetCmd)
}
