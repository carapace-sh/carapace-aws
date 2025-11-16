package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getCaseAuditEventsCmd = &cobra.Command{
	Use:   "get-case-audit-events",
	Short: "Returns the audit history about a specific case if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getCaseAuditEventsCmd).Standalone()

	connectcases_getCaseAuditEventsCmd.Flags().String("case-id", "", "A unique identifier of the case.")
	connectcases_getCaseAuditEventsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_getCaseAuditEventsCmd.Flags().String("max-results", "", "The maximum number of audit events to return.")
	connectcases_getCaseAuditEventsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_getCaseAuditEventsCmd.MarkFlagRequired("case-id")
	connectcases_getCaseAuditEventsCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_getCaseAuditEventsCmd)
}
