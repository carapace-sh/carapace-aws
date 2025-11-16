package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteSpotDatafeedSubscriptionCmd = &cobra.Command{
	Use:   "delete-spot-datafeed-subscription",
	Short: "Deletes the data feed for Spot Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteSpotDatafeedSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteSpotDatafeedSubscriptionCmd).Standalone()

		ec2_deleteSpotDatafeedSubscriptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteSpotDatafeedSubscriptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteSpotDatafeedSubscriptionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteSpotDatafeedSubscriptionCmd)
}
