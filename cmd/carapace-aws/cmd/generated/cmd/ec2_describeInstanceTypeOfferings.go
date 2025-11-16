package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceTypeOfferingsCmd = &cobra.Command{
	Use:   "describe-instance-type-offerings",
	Short: "Lists the instance types that are offered for the specified location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceTypeOfferingsCmd).Standalone()

	ec2_describeInstanceTypeOfferingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceTypeOfferingsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeInstanceTypeOfferingsCmd.Flags().String("location-type", "", "The location type.")
	ec2_describeInstanceTypeOfferingsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceTypeOfferingsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceTypeOfferingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceTypeOfferingsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceTypeOfferingsCmd)
}
