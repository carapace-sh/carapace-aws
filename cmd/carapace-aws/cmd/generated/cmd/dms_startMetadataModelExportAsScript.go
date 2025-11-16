package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelExportAsScriptCmd = &cobra.Command{
	Use:   "start-metadata-model-export-as-script",
	Short: "Saves your converted code to a file as a SQL script, and stores this file on your Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelExportAsScriptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startMetadataModelExportAsScriptCmd).Standalone()

		dms_startMetadataModelExportAsScriptCmd.Flags().String("file-name", "", "The name of the model file to create in the Amazon S3 bucket.")
		dms_startMetadataModelExportAsScriptCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_startMetadataModelExportAsScriptCmd.Flags().String("origin", "", "Whether to export the metadata model from the source or the target.")
		dms_startMetadataModelExportAsScriptCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to export.")
		dms_startMetadataModelExportAsScriptCmd.MarkFlagRequired("migration-project-identifier")
		dms_startMetadataModelExportAsScriptCmd.MarkFlagRequired("origin")
		dms_startMetadataModelExportAsScriptCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_startMetadataModelExportAsScriptCmd)
}
