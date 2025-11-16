package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchUpdateFindingsV2Cmd = &cobra.Command{
	Use:   "batch-update-findings-v2",
	Short: "Used by customers to update information about their investigation into a finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchUpdateFindingsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchUpdateFindingsV2Cmd).Standalone()

		securityhub_batchUpdateFindingsV2Cmd.Flags().String("comment", "", "The updated value for a user provided comment about the finding.")
		securityhub_batchUpdateFindingsV2Cmd.Flags().String("finding-identifiers", "", "Provides information to identify a specific V2 finding.")
		securityhub_batchUpdateFindingsV2Cmd.Flags().String("metadata-uids", "", "The list of finding `metadata.uid` to indicate findings to update.")
		securityhub_batchUpdateFindingsV2Cmd.Flags().String("severity-id", "", "The updated value for the normalized severity identifier.")
		securityhub_batchUpdateFindingsV2Cmd.Flags().String("status-id", "", "The updated value for the normalized status identifier.")
	})
	securityhubCmd.AddCommand(securityhub_batchUpdateFindingsV2Cmd)
}
