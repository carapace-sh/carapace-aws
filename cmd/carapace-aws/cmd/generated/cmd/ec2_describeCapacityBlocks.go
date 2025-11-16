package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityBlocksCmd = &cobra.Command{
	Use:   "describe-capacity-blocks",
	Short: "Describes details about Capacity Blocks in the Amazon Web Services Region that you're currently using.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityBlocksCmd).Standalone()

	ec2_describeCapacityBlocksCmd.Flags().String("capacity-block-ids", "", "The IDs of the Capacity Blocks.")
	ec2_describeCapacityBlocksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityBlocksCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeCapacityBlocksCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeCapacityBlocksCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2_describeCapacityBlocksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityBlocksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeCapacityBlocksCmd)
}
