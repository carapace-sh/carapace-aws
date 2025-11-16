package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityBlockOfferingsCmd = &cobra.Command{
	Use:   "describe-capacity-block-offerings",
	Short: "Describes Capacity Block offerings available for purchase in the Amazon Web Services Region that you're currently using.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityBlockOfferingsCmd).Standalone()

	ec2_describeCapacityBlockOfferingsCmd.Flags().String("capacity-duration-hours", "", "The reservation duration for the Capacity Block, in hours.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("end-date-range", "", "The latest end date for the Capacity Block offering.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("instance-count", "", "The number of instances for which to reserve capacity.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("instance-type", "", "The type of instance for which the Capacity Block offering reserves capacity.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("start-date-range", "", "The earliest start date for the Capacity Block offering.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("ultraserver-count", "", "The number of EC2 UltraServers in the offerings.")
	ec2_describeCapacityBlockOfferingsCmd.Flags().String("ultraserver-type", "", "The EC2 UltraServer type of the Capacity Block offerings.")
	ec2_describeCapacityBlockOfferingsCmd.MarkFlagRequired("capacity-duration-hours")
	ec2_describeCapacityBlockOfferingsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeCapacityBlockOfferingsCmd)
}
