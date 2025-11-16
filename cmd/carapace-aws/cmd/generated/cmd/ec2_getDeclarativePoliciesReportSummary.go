package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getDeclarativePoliciesReportSummaryCmd = &cobra.Command{
	Use:   "get-declarative-policies-report-summary",
	Short: "Retrieves a summary of the account status report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getDeclarativePoliciesReportSummaryCmd).Standalone()

	ec2_getDeclarativePoliciesReportSummaryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getDeclarativePoliciesReportSummaryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getDeclarativePoliciesReportSummaryCmd.Flags().String("report-id", "", "The ID of the report.")
	ec2_getDeclarativePoliciesReportSummaryCmd.Flag("no-dry-run").Hidden = true
	ec2_getDeclarativePoliciesReportSummaryCmd.MarkFlagRequired("report-id")
	ec2Cmd.AddCommand(ec2_getDeclarativePoliciesReportSummaryCmd)
}
