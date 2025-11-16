package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getOpsMetadataCmd = &cobra.Command{
	Use:   "get-ops-metadata",
	Short: "View operational metadata related to an application in Application Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getOpsMetadataCmd).Standalone()

	ssm_getOpsMetadataCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getOpsMetadataCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_getOpsMetadataCmd.Flags().String("ops-metadata-arn", "", "The Amazon Resource Name (ARN) of an OpsMetadata Object to view.")
	ssm_getOpsMetadataCmd.MarkFlagRequired("ops-metadata-arn")
	ssmCmd.AddCommand(ssm_getOpsMetadataCmd)
}
