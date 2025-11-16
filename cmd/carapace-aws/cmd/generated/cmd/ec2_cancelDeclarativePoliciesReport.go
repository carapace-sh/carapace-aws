package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelDeclarativePoliciesReportCmd = &cobra.Command{
	Use:   "cancel-declarative-policies-report",
	Short: "Cancels the generation of an account status report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelDeclarativePoliciesReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelDeclarativePoliciesReportCmd).Standalone()

		ec2_cancelDeclarativePoliciesReportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelDeclarativePoliciesReportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelDeclarativePoliciesReportCmd.Flags().String("report-id", "", "The ID of the report.")
		ec2_cancelDeclarativePoliciesReportCmd.Flag("no-dry-run").Hidden = true
		ec2_cancelDeclarativePoliciesReportCmd.MarkFlagRequired("report-id")
	})
	ec2Cmd.AddCommand(ec2_cancelDeclarativePoliciesReportCmd)
}
