package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstancesCmd = &cobra.Command{
	Use:   "describe-instances",
	Short: "Describes the specified instances or all instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeInstancesCmd).Standalone()

		ec2_describeInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeInstancesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeInstancesCmd.Flags().String("instance-ids", "", "The instance IDs.")
		ec2_describeInstancesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeInstancesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeInstancesCmd)
}
