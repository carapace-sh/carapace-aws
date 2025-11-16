package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityBlockExtensionOfferingsCmd = &cobra.Command{
	Use:   "describe-capacity-block-extension-offerings",
	Short: "Describes Capacity Block extension offerings available for purchase in the Amazon Web Services Region that you're currently using.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityBlockExtensionOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityBlockExtensionOfferingsCmd).Standalone()

		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().String("capacity-block-extension-duration-hours", "", "The duration of the Capacity Block extension offering in hours.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity reservation to be extended.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityBlockExtensionOfferingsCmd.MarkFlagRequired("capacity-block-extension-duration-hours")
		ec2_describeCapacityBlockExtensionOfferingsCmd.MarkFlagRequired("capacity-reservation-id")
		ec2_describeCapacityBlockExtensionOfferingsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityBlockExtensionOfferingsCmd)
}
