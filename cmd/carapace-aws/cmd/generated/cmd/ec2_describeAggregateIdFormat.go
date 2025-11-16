package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAggregateIdFormatCmd = &cobra.Command{
	Use:   "describe-aggregate-id-format",
	Short: "Describes the longer ID format settings for all resource types in a specific Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAggregateIdFormatCmd).Standalone()

	ec2_describeAggregateIdFormatCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAggregateIdFormatCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAggregateIdFormatCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeAggregateIdFormatCmd)
}
