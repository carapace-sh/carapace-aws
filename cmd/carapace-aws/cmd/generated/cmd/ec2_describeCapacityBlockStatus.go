package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityBlockStatusCmd = &cobra.Command{
	Use:   "describe-capacity-block-status",
	Short: "Describes the availability of capacity for the specified Capacity blocks, or all of your Capacity Blocks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityBlockStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityBlockStatusCmd).Standalone()

		ec2_describeCapacityBlockStatusCmd.Flags().String("capacity-block-ids", "", "The ID of the Capacity Block.")
		ec2_describeCapacityBlockStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockStatusCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeCapacityBlockStatusCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityBlockStatusCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeCapacityBlockStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockStatusCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityBlockStatusCmd)
}
