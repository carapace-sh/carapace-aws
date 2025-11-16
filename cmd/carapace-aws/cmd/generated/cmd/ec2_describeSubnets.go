package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSubnetsCmd = &cobra.Command{
	Use:   "describe-subnets",
	Short: "Describes your subnets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSubnetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSubnetsCmd).Standalone()

		ec2_describeSubnetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSubnetsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeSubnetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSubnetsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeSubnetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSubnetsCmd.Flags().String("subnet-ids", "", "The IDs of the subnets.")
		ec2_describeSubnetsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSubnetsCmd)
}
