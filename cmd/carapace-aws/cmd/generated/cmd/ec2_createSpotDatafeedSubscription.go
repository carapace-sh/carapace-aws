package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSpotDatafeedSubscriptionCmd = &cobra.Command{
	Use:   "create-spot-datafeed-subscription",
	Short: "Creates a data feed for Spot Instances, enabling you to view Spot Instance usage logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSpotDatafeedSubscriptionCmd).Standalone()

	ec2_createSpotDatafeedSubscriptionCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket in which to store the Spot Instance data feed.")
	ec2_createSpotDatafeedSubscriptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSpotDatafeedSubscriptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSpotDatafeedSubscriptionCmd.Flags().String("prefix", "", "The prefix for the data feed file names.")
	ec2_createSpotDatafeedSubscriptionCmd.MarkFlagRequired("bucket")
	ec2_createSpotDatafeedSubscriptionCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createSpotDatafeedSubscriptionCmd)
}
