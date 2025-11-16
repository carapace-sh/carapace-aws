package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listOpsMetadataCmd = &cobra.Command{
	Use:   "list-ops-metadata",
	Short: "Amazon Web Services Systems Manager calls this API operation when displaying all Application Manager OpsMetadata objects or blobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listOpsMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listOpsMetadataCmd).Standalone()

		ssm_listOpsMetadataCmd.Flags().String("filters", "", "One or more filters to limit the number of OpsMetadata objects returned by the call.")
		ssm_listOpsMetadataCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listOpsMetadataCmd.Flags().String("next-token", "", "A token to start the list.")
	})
	ssmCmd.AddCommand(ssm_listOpsMetadataCmd)
}
