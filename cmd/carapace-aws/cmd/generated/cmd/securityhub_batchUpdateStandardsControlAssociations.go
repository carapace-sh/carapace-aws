package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchUpdateStandardsControlAssociationsCmd = &cobra.Command{
	Use:   "batch-update-standards-control-associations",
	Short: "For a batch of security controls and standards, this operation updates the enablement status of a control in a standard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchUpdateStandardsControlAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchUpdateStandardsControlAssociationsCmd).Standalone()

		securityhub_batchUpdateStandardsControlAssociationsCmd.Flags().String("standards-control-association-updates", "", "Updates the enablement status of a security control in a specified standard.")
		securityhub_batchUpdateStandardsControlAssociationsCmd.MarkFlagRequired("standards-control-association-updates")
	})
	securityhubCmd.AddCommand(securityhub_batchUpdateStandardsControlAssociationsCmd)
}
