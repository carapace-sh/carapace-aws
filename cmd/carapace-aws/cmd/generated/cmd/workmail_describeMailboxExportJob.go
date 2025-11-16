package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeMailboxExportJobCmd = &cobra.Command{
	Use:   "describe-mailbox-export-job",
	Short: "Describes the current status of a mailbox export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeMailboxExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_describeMailboxExportJobCmd).Standalone()

		workmail_describeMailboxExportJobCmd.Flags().String("job-id", "", "The mailbox export job ID.")
		workmail_describeMailboxExportJobCmd.Flags().String("organization-id", "", "The organization ID.")
		workmail_describeMailboxExportJobCmd.MarkFlagRequired("job-id")
		workmail_describeMailboxExportJobCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_describeMailboxExportJobCmd)
}
