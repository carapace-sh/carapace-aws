package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_startDeclarativePoliciesReportCmd = &cobra.Command{
	Use:   "start-declarative-policies-report",
	Short: "Generates an account status report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_startDeclarativePoliciesReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_startDeclarativePoliciesReportCmd).Standalone()

		ec2_startDeclarativePoliciesReportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_startDeclarativePoliciesReportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_startDeclarativePoliciesReportCmd.Flags().String("s3-bucket", "", "The name of the S3 bucket where the report will be saved.")
		ec2_startDeclarativePoliciesReportCmd.Flags().String("s3-prefix", "", "The prefix for your S3 object.")
		ec2_startDeclarativePoliciesReportCmd.Flags().String("tag-specifications", "", "The tags to apply.")
		ec2_startDeclarativePoliciesReportCmd.Flags().String("target-id", "", "The root ID, organizational unit ID, or account ID.")
		ec2_startDeclarativePoliciesReportCmd.Flag("no-dry-run").Hidden = true
		ec2_startDeclarativePoliciesReportCmd.MarkFlagRequired("s3-bucket")
		ec2_startDeclarativePoliciesReportCmd.MarkFlagRequired("target-id")
	})
	ec2Cmd.AddCommand(ec2_startDeclarativePoliciesReportCmd)
}
