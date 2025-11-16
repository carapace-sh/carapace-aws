package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listMailboxExportJobsCmd = &cobra.Command{
	Use:   "list-mailbox-export-jobs",
	Short: "Lists the mailbox export jobs started for the specified organization within the last seven days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listMailboxExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listMailboxExportJobsCmd).Standalone()

		workmail_listMailboxExportJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listMailboxExportJobsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listMailboxExportJobsCmd.Flags().String("organization-id", "", "The organization ID.")
		workmail_listMailboxExportJobsCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listMailboxExportJobsCmd)
}
