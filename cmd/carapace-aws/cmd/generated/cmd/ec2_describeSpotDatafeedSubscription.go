package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotDatafeedSubscriptionCmd = &cobra.Command{
	Use:   "describe-spot-datafeed-subscription",
	Short: "Describes the data feed for Spot Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotDatafeedSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSpotDatafeedSubscriptionCmd).Standalone()

		ec2_describeSpotDatafeedSubscriptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotDatafeedSubscriptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotDatafeedSubscriptionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSpotDatafeedSubscriptionCmd)
}
