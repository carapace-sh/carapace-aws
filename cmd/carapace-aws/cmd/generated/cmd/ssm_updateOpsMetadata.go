package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateOpsMetadataCmd = &cobra.Command{
	Use:   "update-ops-metadata",
	Short: "Amazon Web Services Systems Manager calls this API operation when you edit OpsMetadata in Application Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateOpsMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateOpsMetadataCmd).Standalone()

		ssm_updateOpsMetadataCmd.Flags().String("keys-to-delete", "", "The metadata keys to delete from the OpsMetadata object.")
		ssm_updateOpsMetadataCmd.Flags().String("metadata-to-update", "", "Metadata to add to an OpsMetadata object.")
		ssm_updateOpsMetadataCmd.Flags().String("ops-metadata-arn", "", "The Amazon Resource Name (ARN) of the OpsMetadata Object to update.")
		ssm_updateOpsMetadataCmd.MarkFlagRequired("ops-metadata-arn")
	})
	ssmCmd.AddCommand(ssm_updateOpsMetadataCmd)
}
