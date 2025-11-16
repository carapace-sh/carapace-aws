package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelConversionCmd = &cobra.Command{
	Use:   "start-metadata-model-conversion",
	Short: "Converts your source database objects to a format compatible with the target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelConversionCmd).Standalone()

	dms_startMetadataModelConversionCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_startMetadataModelConversionCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to convert.")
	dms_startMetadataModelConversionCmd.MarkFlagRequired("migration-project-identifier")
	dms_startMetadataModelConversionCmd.MarkFlagRequired("selection-rules")
	dmsCmd.AddCommand(dms_startMetadataModelConversionCmd)
}
