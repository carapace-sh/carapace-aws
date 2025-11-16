package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuditFindingsCmd = &cobra.Command{
	Use:   "list-audit-findings",
	Short: "Lists the findings (results) of a Device Defender audit or of the audits performed during a specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuditFindingsCmd).Standalone()

	iot_listAuditFindingsCmd.Flags().String("check-name", "", "A filter to limit results to the findings for the specified audit check.")
	iot_listAuditFindingsCmd.Flags().String("end-time", "", "A filter to limit results to those found before the specified time.")
	iot_listAuditFindingsCmd.Flags().String("list-suppressed-findings", "", "Boolean flag indicating whether only the suppressed findings or the unsuppressed findings should be listed.")
	iot_listAuditFindingsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listAuditFindingsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listAuditFindingsCmd.Flags().String("resource-identifier", "", "Information identifying the noncompliant resource.")
	iot_listAuditFindingsCmd.Flags().String("start-time", "", "A filter to limit results to those found after the specified time.")
	iot_listAuditFindingsCmd.Flags().String("task-id", "", "A filter to limit results to the audit with the specified ID.")
	iotCmd.AddCommand(iot_listAuditFindingsCmd)
}
