package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_cancelMetadataModelCreationCmd = &cobra.Command{
	Use:   "cancel-metadata-model-creation",
	Short: "Cancels a single metadata model creation operation that was started with `StartMetadataModelCreation`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_cancelMetadataModelCreationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_cancelMetadataModelCreationCmd).Standalone()

		dms_cancelMetadataModelCreationCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_cancelMetadataModelCreationCmd.Flags().String("request-identifier", "", "The identifier for the metadata model creation operation to cancel.")
		dms_cancelMetadataModelCreationCmd.MarkFlagRequired("migration-project-identifier")
		dms_cancelMetadataModelCreationCmd.MarkFlagRequired("request-identifier")
	})
	dmsCmd.AddCommand(dms_cancelMetadataModelCreationCmd)
}
