package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceStatusCmd = &cobra.Command{
	Use:   "describe-instance-status",
	Short: "Describes the status of the specified instances or all of your instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceStatusCmd).Standalone()

	ec2_describeInstanceStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceStatusCmd.Flags().String("filters", "", "The filters.")
	ec2_describeInstanceStatusCmd.Flags().Bool("include-all-instances", false, "When `true`, includes the health status for all instances.")
	ec2_describeInstanceStatusCmd.Flags().String("instance-ids", "", "The instance IDs.")
	ec2_describeInstanceStatusCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceStatusCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceStatusCmd.Flags().Bool("no-include-all-instances", false, "When `true`, includes the health status for all instances.")
	ec2_describeInstanceStatusCmd.Flag("no-dry-run").Hidden = true
	ec2_describeInstanceStatusCmd.Flag("no-include-all-instances").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceStatusCmd)
}
