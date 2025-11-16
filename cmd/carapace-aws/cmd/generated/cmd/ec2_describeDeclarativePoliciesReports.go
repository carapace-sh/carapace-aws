package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeDeclarativePoliciesReportsCmd = &cobra.Command{
	Use:   "describe-declarative-policies-reports",
	Short: "Describes the metadata of an account status report, including the status of the report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeDeclarativePoliciesReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeDeclarativePoliciesReportsCmd).Standalone()

		ec2_describeDeclarativePoliciesReportsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeDeclarativePoliciesReportsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeDeclarativePoliciesReportsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeDeclarativePoliciesReportsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeDeclarativePoliciesReportsCmd.Flags().String("report-ids", "", "One or more report IDs.")
		ec2_describeDeclarativePoliciesReportsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeDeclarativePoliciesReportsCmd)
}
