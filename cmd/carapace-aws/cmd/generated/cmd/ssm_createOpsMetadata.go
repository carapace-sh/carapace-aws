package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createOpsMetadataCmd = &cobra.Command{
	Use:   "create-ops-metadata",
	Short: "If you create a new application in Application Manager, Amazon Web Services Systems Manager calls this API operation to specify information about the new application, including the application type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createOpsMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_createOpsMetadataCmd).Standalone()

		ssm_createOpsMetadataCmd.Flags().String("metadata", "", "Metadata for a new Application Manager application.")
		ssm_createOpsMetadataCmd.Flags().String("resource-id", "", "A resource ID for a new Application Manager application.")
		ssm_createOpsMetadataCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_createOpsMetadataCmd.MarkFlagRequired("resource-id")
	})
	ssmCmd.AddCommand(ssm_createOpsMetadataCmd)
}
