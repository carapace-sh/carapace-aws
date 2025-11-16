package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteOpsMetadataCmd = &cobra.Command{
	Use:   "delete-ops-metadata",
	Short: "Delete OpsMetadata related to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteOpsMetadataCmd).Standalone()

	ssm_deleteOpsMetadataCmd.Flags().String("ops-metadata-arn", "", "The Amazon Resource Name (ARN) of an OpsMetadata Object to delete.")
	ssm_deleteOpsMetadataCmd.MarkFlagRequired("ops-metadata-arn")
	ssmCmd.AddCommand(ssm_deleteOpsMetadataCmd)
}
