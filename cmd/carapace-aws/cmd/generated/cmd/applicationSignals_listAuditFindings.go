package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listAuditFindingsCmd = &cobra.Command{
	Use:   "list-audit-findings",
	Short: "Retrieves a list of audit findings for Application Signals resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listAuditFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listAuditFindingsCmd).Standalone()

		applicationSignals_listAuditFindingsCmd.Flags().String("audit-targets", "", "An array of audit target specifications to filter the findings.")
		applicationSignals_listAuditFindingsCmd.Flags().String("auditors", "", "An array of auditor names to filter the findings.")
		applicationSignals_listAuditFindingsCmd.Flags().String("end-time", "", "The end time for the audit findings query.")
		applicationSignals_listAuditFindingsCmd.Flags().String("max-results", "", "The maximum number of audit findings to return in a single request.")
		applicationSignals_listAuditFindingsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		applicationSignals_listAuditFindingsCmd.Flags().String("start-time", "", "The start time for the audit findings query.")
		applicationSignals_listAuditFindingsCmd.MarkFlagRequired("audit-targets")
		applicationSignals_listAuditFindingsCmd.MarkFlagRequired("end-time")
		applicationSignals_listAuditFindingsCmd.MarkFlagRequired("start-time")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listAuditFindingsCmd)
}
