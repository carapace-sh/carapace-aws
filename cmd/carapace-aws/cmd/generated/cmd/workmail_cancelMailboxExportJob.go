package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_cancelMailboxExportJobCmd = &cobra.Command{
	Use:   "cancel-mailbox-export-job",
	Short: "Cancels a mailbox export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_cancelMailboxExportJobCmd).Standalone()

	workmail_cancelMailboxExportJobCmd.Flags().String("client-token", "", "The idempotency token for the client request.")
	workmail_cancelMailboxExportJobCmd.Flags().String("job-id", "", "The job ID.")
	workmail_cancelMailboxExportJobCmd.Flags().String("organization-id", "", "The organization ID.")
	workmail_cancelMailboxExportJobCmd.MarkFlagRequired("client-token")
	workmail_cancelMailboxExportJobCmd.MarkFlagRequired("job-id")
	workmail_cancelMailboxExportJobCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_cancelMailboxExportJobCmd)
}
