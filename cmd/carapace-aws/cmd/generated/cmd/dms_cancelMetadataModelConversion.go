package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_cancelMetadataModelConversionCmd = &cobra.Command{
	Use:   "cancel-metadata-model-conversion",
	Short: "Cancels a single metadata model conversion operation that was started with `StartMetadataModelConversion`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_cancelMetadataModelConversionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_cancelMetadataModelConversionCmd).Standalone()

		dms_cancelMetadataModelConversionCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_cancelMetadataModelConversionCmd.Flags().String("request-identifier", "", "The identifier for the metadata model conversion operation to cancel.")
		dms_cancelMetadataModelConversionCmd.MarkFlagRequired("migration-project-identifier")
		dms_cancelMetadataModelConversionCmd.MarkFlagRequired("request-identifier")
	})
	dmsCmd.AddCommand(dms_cancelMetadataModelConversionCmd)
}
