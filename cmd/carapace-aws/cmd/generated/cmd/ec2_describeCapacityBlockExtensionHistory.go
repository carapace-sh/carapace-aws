package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityBlockExtensionHistoryCmd = &cobra.Command{
	Use:   "describe-capacity-block-extension-history",
	Short: "Describes the events for the specified Capacity Block extension during the specified time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityBlockExtensionHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityBlockExtensionHistoryCmd).Standalone()

		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().String("capacity-reservation-ids", "", "The IDs of Capacity Block reservations that you want to display the history for.")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().String("filters", "", "One or more filters")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockExtensionHistoryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityBlockExtensionHistoryCmd)
}
