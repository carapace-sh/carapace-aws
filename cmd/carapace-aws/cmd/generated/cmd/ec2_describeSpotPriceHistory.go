package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotPriceHistoryCmd = &cobra.Command{
	Use:   "describe-spot-price-history",
	Short: "Describes the Spot price history.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotPriceHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSpotPriceHistoryCmd).Standalone()

		ec2_describeSpotPriceHistoryCmd.Flags().String("availability-zone", "", "Filters the results by the specified Availability Zone.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("availability-zone-id", "", "Filters the results by the specified ID of the Availability Zone.")
		ec2_describeSpotPriceHistoryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("end-time", "", "The date and time, up to the current date, from which to stop retrieving the price history data, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
		ec2_describeSpotPriceHistoryCmd.Flags().String("filters", "", "The filters.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("instance-types", "", "Filters the results by the specified instance types.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeSpotPriceHistoryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("product-descriptions", "", "Filters the results by the specified basic product descriptions.")
		ec2_describeSpotPriceHistoryCmd.Flags().String("start-time", "", "The date and time, up to the past 90 days, from which to start retrieving the price history data, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
		ec2_describeSpotPriceHistoryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSpotPriceHistoryCmd)
}
